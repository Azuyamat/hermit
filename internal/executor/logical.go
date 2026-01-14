package executor

import (
	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/token"
)

func (e *Executor) executeLogicalExpr(expr *ast.LogicalExpr) error {
	leftErr := e.executeStatement(expr.Left)

	if expr.Operator == token.AND {
		if leftErr != nil {
			return leftErr
		}
		return e.executeStatement(expr.Right)
	}

	if expr.Operator == token.OR {
		if leftErr == nil {
			return nil
		}
		return e.executeStatement(expr.Right)
	}

	return nil
}
