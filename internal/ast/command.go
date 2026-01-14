package ast

import (
	"strings"

	"github.com/azuyamat/hermit/internal/token"
)

type Command struct {
	Token     token.Token
	Name      string
	Args      []Argument
	Redirects []Redirect
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
		out.WriteString(arg.String())
	}

	for _, redirect := range c.Redirects {
		out.WriteString(" ")
		out.WriteString(redirect.String())
	}

	return out.String()
}
