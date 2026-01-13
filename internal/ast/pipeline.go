package ast

import (
	"strings"

	"github.com/azuyamat/hermit/internal/token"
)

type Pipeline struct {
	Token    token.Token
	Commands []*Command
}

func (p *Pipeline) statementNode() {}

func (p *Pipeline) TokenLiteral() string {
	return p.Token.Literal
}

func (p *Pipeline) String() string {
	var out strings.Builder
	for i, cmd := range p.Commands {
		out.WriteString(cmd.String())
		if i < len(p.Commands)-1 {
			out.WriteString(" | ")
		}
	}
	return out.String()
}
