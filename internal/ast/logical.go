package ast

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/token"
)

type LogicalExpr struct {
	Token    token.Token
	Left     Statement
	Operator token.TokenType
	Right    Statement
}

func (lo *LogicalExpr) statementNode() {}

func (lo *LogicalExpr) TokenLiteral() string {
	return lo.Token.Literal
}

func (lo *LogicalExpr) String() string {
	return fmt.Sprintf("%s %s %s", lo.Left.String(), lo.Operator, lo.Right.String())
}
