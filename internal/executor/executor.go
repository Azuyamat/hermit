package executor

import (
	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/builtins"
	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Executor struct {
	builtins *command.Manager
	context  *types.ExecutionContext
}

func New() *Executor {
	cmdManager := command.NewManager()
	builtins.RegisterCoreBuiltins(cmdManager)

	return &Executor{
		builtins: cmdManager,
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
	case *ast.LogicalExpr:
		return e.executeLogicalExpr(s)
	default:
		return nil
	}
}

func (e *Executor) Context() *types.ExecutionContext {
	return e.context
}
