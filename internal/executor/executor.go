package executor

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/builtins"
	"github.com/azuyamat/hermit/internal/types"
)

type Executor struct {
	builtins *BuiltinRegistry
	context  *types.ExecutionContext
}

func New() *Executor {
	registry := NewBuiltinRegistry()
	registry.Register(builtins.NewCd())
	registry.Register(builtins.NewClear())
	registry.Register(builtins.NewEcho())
	registry.Register(builtins.NewPwd())
	registry.Register(builtins.NewExit())
	registry.Register(builtins.NewExport())
	registry.Register(builtins.NewCat())

	return &Executor{
		builtins: registry,
		context:  types.NewContext(),
	}
}

func (e *Executor) Execute(program *ast.Program) error {
	for _, stmt := range program.Statements {
		if err := e.executeStatement(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (e *Executor) executeStatement(stmt ast.Statement) error {
	switch s := stmt.(type) {
	case *ast.Command:
		return e.executeCommand(s)
	case *ast.Pipeline:
		return e.executePipeline(s)
	default:
		return nil
	}
}

func (e *Executor) executeCommand(cmd *ast.Command) error {
	if builtin, ok := e.builtins.Get(cmd.Name); ok {
		return builtin.Execute(cmd.Args, e.context, e.context.Stdin, e.context.Stdout, e.context.Stderr)
	}

	command := exec.CommandContext(e.context.Context, cmd.Name, cmd.Args...)
	command.Stdin = e.context.Stdin
	command.Stdout = e.context.Stdout
	command.Stderr = e.context.Stderr
	command.Env = e.context.EnvSlice()

	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			e.context.LastExitCode = exitError.ExitCode()
		} else {
			e.context.LastExitCode = 1
		}
	} else {
		e.context.LastExitCode = 0
	}

	return err
}

func (e *Executor) executePipeline(pipeline *ast.Pipeline) error {
	if len(pipeline.Commands) == 0 {
		return fmt.Errorf("empty pipeline")
	}

	if len(pipeline.Commands) == 1 {
		return e.executeCommand(pipeline.Commands[0])
	}

	numCommands := len(pipeline.Commands)
	pipes := make([]io.ReadCloser, numCommands-1)

	for i := 0; i < numCommands-1; i++ {
		r, w := io.Pipe()
		pipes[i] = r

		go func(cmdIndex int, stdin io.Reader, stdout io.WriteCloser) {
			defer stdout.Close()

			astCmd := pipeline.Commands[cmdIndex]

			if builtin, ok := e.builtins.Get(astCmd.Name); ok {
				builtin.Execute(astCmd.Args, e.context, stdin, stdout, e.context.Stderr)
			} else {
				cmd := exec.CommandContext(e.context.Context, astCmd.Name, astCmd.Args...)
				cmd.Stdin = stdin
				cmd.Stdout = stdout
				cmd.Stderr = e.context.Stderr
				cmd.Env = e.context.EnvSlice()
				cmd.Run()
			}
		}(i, e.getStdinForCommand(i, pipes), w)
	}

	lastCmd := pipeline.Commands[numCommands-1]
	lastStdin := e.getStdinForCommand(numCommands-1, pipes)

	if builtin, ok := e.builtins.Get(lastCmd.Name); ok {
		err := builtin.Execute(lastCmd.Args, e.context, lastStdin, e.context.Stdout, e.context.Stderr)
		if err != nil {
			e.context.LastExitCode = 1
			return err
		}
		e.context.LastExitCode = 0
		return nil
	}

	cmd := exec.CommandContext(e.context.Context, lastCmd.Name, lastCmd.Args...)
	cmd.Stdin = lastStdin
	cmd.Stdout = e.context.Stdout
	cmd.Stderr = e.context.Stderr
	cmd.Env = e.context.EnvSlice()

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			e.context.LastExitCode = exitError.ExitCode()
		} else {
			e.context.LastExitCode = 1
		}
		return err
	}

	e.context.LastExitCode = 0
	return nil
}

func (e *Executor) getStdinForCommand(index int, pipes []io.ReadCloser) io.Reader {
	if index == 0 {
		return e.context.Stdin
	}
	return pipes[index-1]
}

func (e *Executor) Context() *types.ExecutionContext {
	return e.context
}
