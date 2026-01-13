package parser

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/token"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	return p.parsePipeline()
}

func (p *Parser) parseCommand() (*ast.Command, error) {
	if !p.curTokenIs(token.IDENT) {
		return nil, fmt.Errorf("expected command name, got %s", p.currentToken.Literal)
	}

	cmd := &ast.Command{
		Name: p.currentToken.Literal,
		Args: []string{},
	}

	p.nextToken()

	for p.curTokenIs(token.IDENT) ||
		p.curTokenIs(token.FLAG) ||
		p.curTokenIs(token.DOUBLE_QUOTED_STRING) ||
		p.curTokenIs(token.SINGLE_QUOTED_STRING) ||
		p.curTokenIs(token.DOT) ||
		p.curTokenIs(token.INT) {
		cmd.Args = append(cmd.Args, p.currentToken.Literal)
		p.nextToken()
	}

	return cmd, nil
}

func (p *Parser) parsePipeline() (ast.Statement, error) {
	left, err := p.parseCommand()
	if err != nil {
		return nil, err
	}

	if !p.curTokenIs(token.PIPE) {
		return left, nil
	}

	commands := []*ast.Command{left}

	for p.curTokenIs(token.PIPE) {
		p.nextToken() // consume |

		right, err := p.parseCommand()
		if err != nil {
			return nil, fmt.Errorf("expected command after pipe: %w", err)
		}

		commands = append(commands, right)
	}

	return &ast.Pipeline{
		Token:    left.Token,
		Commands: commands,
	}, nil
}
