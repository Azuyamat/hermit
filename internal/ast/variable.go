package ast

import "github.com/azuyamat/hermit/internal/token"

type Variable struct {
	Token token.Token
	Name  string
}

func (v *Variable) argumentNode() {}

func (v *Variable) TokenLiteral() string {
	return v.Token.Literal
}

func (v *Variable) String() string {
	return "$" + v.Name
}
