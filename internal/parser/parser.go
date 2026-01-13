package parser

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/token"
)

type Lexer interface {
	NextToken() (token.Token, error)
}

type Parser struct {
	lexer        Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []error
}

func New(lexer Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{}

	for !p.curTokenIs(token.EOF) {
		statement, err := p.parseStatement()
		if err != nil {
			p.errors = append(p.errors, err)
			p.displayErrors()
			return nil, fmt.Errorf("failed to parse statement: %w", err)
		}
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	if len(p.errors) > 0 {
		p.displayErrors()
		return nil, fmt.Errorf("encountered %d errors while parsing", len(p.errors))
	}

	return program, nil
}

func (p *Parser) displayErrors() {
	for _, msg := range p.errors {
		fmt.Printf("\t%s\n", msg)
	}
	fmt.Println()
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	tok, err := p.lexer.NextToken()
	if err != nil {
		p.errors = append(p.errors, err)
		return
	}
	p.peekToken = tok
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) Errors() []error {
	return p.errors
}
