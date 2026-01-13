package ast

import (
	"strings"

	"github.com/azuyamat/hermit/internal/token"
)

type Command struct {
	Token token.Token
	Name  string
	Args  []string
}

func (c *Command) statementNode() {}

func (c *Command) TokenLiteral() string {
	return c.Token.Literal
}

func (c *Command) String() string {
	var out strings.Builder
	out.WriteString(c.Name)
	for _, arg := range c.Args {
		out.WriteString(" ")
		out.WriteString(arg)
	}
	return out.String()
}
