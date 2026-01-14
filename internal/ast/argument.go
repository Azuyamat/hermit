package ast

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/token"
)

type Argument interface {
	Node
	argumentNode()
}

type LiteralArg struct {
	Token token.Token
	Value string
}

func (la *LiteralArg) argumentNode() {}

func (la *LiteralArg) TokenLiteral() string {
	return la.Token.Literal
}

func (la *LiteralArg) String() string {
	return la.Token.Literal
}

type QuotedString struct {
	Token token.Token
	Value string
	Quote token.TokenType
}

func (qs *QuotedString) argumentNode() {}

func (qs *QuotedString) TokenLiteral() string {
	return qs.Token.Literal
}

func (qs *QuotedString) String() string {
	return qs.Value
}

type CommandSubstitution struct {
	Token     token.Token
	Statement Statement
}

func (cs *CommandSubstitution) argumentNode() {}

func (cs *CommandSubstitution) TokenLiteral() string {
	return cs.Token.Literal
}

func (cs *CommandSubstitution) String() string {
	return fmt.Sprintf("$(%s)", cs.Statement.String())
}
