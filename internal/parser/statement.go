package parser

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/token"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	return p.parseLogicalOr()
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

func (p *Parser) parseCommand() (*ast.Command, error) {
	if !p.curTokenIs(token.IDENT) {
		return nil, fmt.Errorf("expected command name, got %s", p.currentToken.Literal)
	}

	cmd := &ast.Command{
		Token:     p.currentToken,
		Name:      p.currentToken.Literal,
		Args:      []ast.Argument{},
		Redirects: []ast.Redirect{},
	}

	p.nextToken()

	for p.isArgumentToken() || p.isRedirectToken() {
		if p.isRedirectToken() {
			redirect, err := p.parseRedirect()
			if err != nil {
				return nil, err
			}
			cmd.Redirects = append(cmd.Redirects, redirect)
		} else {
			arg, err := p.parseArgument()
			if err != nil {
				return nil, err
			}
			cmd.Args = append(cmd.Args, arg)
		}
		p.nextToken()
	}

	return cmd, nil
}

func (p *Parser) parseRedirect() (ast.Redirect, error) {
	redirectToken := p.currentToken

	var redirectType ast.RedirectType
	switch p.currentToken.Type {
	case token.REDIRECT_OUT:
		redirectType = ast.RedirectStdout
	case token.REDIRECT_APPEND:
		redirectType = ast.RedirectAppend
	case token.REDIRECT_IN:
		redirectType = ast.RedirectStdin
	case token.REDIRECT_ERR:
		redirectType = ast.RedirectStderr
	case token.REDIRECT_ERR_APPEND:
		redirectType = ast.RedirectStderrAppend
	case token.REDIRECT_BOTH:
		redirectType = ast.RedirectBoth
	default:
		return ast.Redirect{}, fmt.Errorf("unexpected redirect token: %s", p.currentToken.Literal)
	}

	p.nextToken()

	if !p.isArgumentToken() {
		return ast.Redirect{}, fmt.Errorf("expected argument after redirect, got %s", p.currentToken.Literal)
	}

	arg, err := p.parseArgument()
	if err != nil {
		return ast.Redirect{}, err
	}

	if _, ok := arg.(*ast.CommandSubstitution); ok {
		return ast.Redirect{}, fmt.Errorf("cannot redirect to command substitution")
	}

	return ast.Redirect{
		Token:  redirectToken,
		Type:   redirectType,
		Target: arg,
	}, nil
}

func (p *Parser) parseArgument() (ast.Argument, error) {
	switch p.currentToken.Type {
	case token.COMMAND_SUB_START:
		return p.parseCommandSubstitution()

	case token.VARIABLE:
		return p.parseVariable()

	case token.IDENT, token.FLAG:
		return &ast.LiteralArg{
			Token: p.currentToken,
			Value: p.currentToken.Literal,
		}, nil

	case token.DOUBLE_QUOTED_STRING, token.SINGLE_QUOTED_STRING:
		return p.parseQuotedString()

	default:
		return nil, fmt.Errorf("unexpected token type for argument: %s", p.currentToken.Type)
	}
}

func (p *Parser) parseCommandSubstitution() (ast.Argument, error) {
	subStart := p.currentToken
	p.nextToken() // consume $(

	innerStatement, err := p.parseStatement()
	if err != nil {
		return nil, fmt.Errorf("failed to parse command substitution: %w", err)
	}

	if !p.curTokenIs(token.RPAREN) {
		return nil, fmt.Errorf("expected closing ), got %s", p.currentToken.Literal)
	}

	return &ast.CommandSubstitution{
		Token:     subStart,
		Statement: innerStatement,
	}, nil
}

func (p *Parser) parseVariable() (ast.Argument, error) {
	varToken := p.currentToken
	name := varToken.Literal[1:]

	if len(name) > 0 && name[0] == '{' && name[len(name)-1] == '}' {
		name = name[1 : len(name)-1]
	}

	return &ast.Variable{
		Token: varToken,
		Name:  name,
	}, nil
}

func (p *Parser) parseQuotedString() (ast.Argument, error) {
	return &ast.QuotedString{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
		Quote: p.currentToken.Type,
	}, nil
}

func (p *Parser) isArgumentToken() bool {
	return p.curTokenIs(token.IDENT) ||
		p.curTokenIs(token.FLAG) ||
		p.curTokenIs(token.DOUBLE_QUOTED_STRING) ||
		p.curTokenIs(token.SINGLE_QUOTED_STRING) ||
		p.curTokenIs(token.COMMAND_SUB_START) ||
		p.curTokenIs(token.VARIABLE)
}

func (p *Parser) isRedirectToken() bool {
	return p.curTokenIs(token.REDIRECT_OUT) ||
		p.curTokenIs(token.REDIRECT_APPEND) ||
		p.curTokenIs(token.REDIRECT_IN) ||
		p.curTokenIs(token.REDIRECT_ERR) ||
		p.curTokenIs(token.REDIRECT_ERR_APPEND) ||
		p.curTokenIs(token.REDIRECT_BOTH)
}
