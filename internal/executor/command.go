package executor

import (
	"errors"
	"fmt"
	"io"
	"os/exec"

	"github.com/azuyamat/hermit/internal/ast"
)

type stdio struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func (e *Executor) executeCommand(cmd *ast.Command) error {
	args, err := e.evaluateArguments(cmd.Args)
	if err != nil {
		return fmt.Errorf("evaluating arguments: %w", err)
	}

	stdio, cleanup, err := e.setupIO(cmd.Redirects)
	if err != nil {
		return fmt.Errorf("setting up I/O: %w", err)
	}
	defer cleanup()

	if _, ok := e.builtins.Get(cmd.Name); ok {
		return e.executeBuiltin(cmd.Name, args, stdio)
	}

	return e.executeExternal(cmd.Name, args, stdio)
}

func (e *Executor) setupIO(redirects []ast.Redirect) (*stdio, func(), error) {
	stdin, stdout, stderr, cleanup, err := e.setupRedirects(redirects)
	if err != nil {
		return nil, nil, err
	}

	io := &stdio{
		stdin:  coalescReader(stdin, e.context.Stdin),
		stdout: coalescWriter(stdout, e.context.Stdout),
		stderr: coalescWriter(stderr, e.context.Stderr),
	}

	return io, cleanup, nil
}

func (e *Executor) executeBuiltin(name string, args []string, stdio *stdio) error {
	err := e.builtins.Execute(name, args, stdio.stdout, stdio.stderr, stdio.stdin, e.context)
	if err != nil {
		e.context.LastExitCode = 1
	} else {
		e.context.LastExitCode = 0
	}
	return err
}

func (e *Executor) executeExternal(name string, args []string, stdio *stdio) error {
	cmd := exec.CommandContext(e.context.Context, name, args...)
	cmd.Stdin = stdio.stdin
	cmd.Stdout = stdio.stdout
	cmd.Stderr = stdio.stderr
	cmd.Env = e.context.EnvSlice()

	err := cmd.Run()

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		e.context.LastExitCode = exitErr.ExitCode()
	} else if err != nil {
		e.context.LastExitCode = 1
	} else {
		e.context.LastExitCode = 0
	}

	if err != nil {
		return fmt.Errorf("command %q failed: %w", name, err)
	}
	return nil
}

func (e *Executor) runCommand(name string, args []string, stdio *stdio) error {
	if _, ok := e.builtins.Get(name); ok {
		return e.executeBuiltin(name, args, stdio)
	}

	cmd := exec.CommandContext(e.context.Context, name, args...)
	cmd.Stdin = stdio.stdin
	cmd.Stdout = stdio.stdout
	cmd.Stderr = stdio.stderr
	cmd.Env = e.context.EnvSlice()
	return cmd.Run()
}
