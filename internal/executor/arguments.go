package executor

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/azuyamat/hermit/internal/ast"
)

func (e *Executor) evaluateArguments(args []ast.Argument) ([]string, error) {
	result := make([]string, 0, len(args))

	for _, arg := range args {
		values, err := e.evaluateArgument(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, values...)
	}

	return result, nil
}

func (e *Executor) evaluateArgument(arg ast.Argument) ([]string, error) {
	switch a := arg.(type) {
	case *ast.LiteralArg:
		return []string{a.Value}, nil
	case *ast.QuotedString:
		return []string{a.Value}, nil
	case *ast.CommandSubstitution:
		return e.evaluateCommandSubstitution(a)
	case *ast.Variable:
		return e.evaluateVariable(a)
	default:
		return nil, fmt.Errorf("unsupported argument type: %T", arg)
	}
}

func (e *Executor) evaluateCommandSubstitution(cs *ast.CommandSubstitution) ([]string, error) {
	var buf bytes.Buffer
	oldStdout := e.context.Stdout
	e.context.Stdout = &buf
	defer func() { e.context.Stdout = oldStdout }()

	if err := e.executeStatement(cs.Statement); err != nil {
		return nil, fmt.Errorf("command substitution: %w", err)
	}

	output := strings.TrimSpace(buf.String())
	if output == "" {
		return []string{}, nil
	}

	return strings.Fields(output), nil
}

func (e *Executor) evaluateVariable(v *ast.Variable) ([]string, error) {
	value, ok := e.context.GetEnv(v.Name)
	if !ok {
		return []string{""}, nil
	}
	return strings.Fields(value), nil
}
