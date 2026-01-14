package parser

import (
	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/token"
)

func (p *Parser) parseLogicalOr() (ast.Statement, error) {
	left, err := p.parseLogicalAnd()
	if err != nil {
		return nil, err
	}

	for p.curTokenIs(token.OR) {
		tok := p.currentToken
		p.nextToken() // consume ||

		right, err := p.parseLogicalAnd()
		if err != nil {
			return nil, err
		}
		left = &ast.LogicalExpr{
			Token:    tok,
			Left:     left,
			Operator: token.OR,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) parseLogicalAnd() (ast.Statement, error) {
	left, err := p.parsePipeline()
	if err != nil {
		return nil, err
	}

	for p.curTokenIs(token.AND) {
		tok := p.currentToken
		p.nextToken() // consume &&

		right, err := p.parsePipeline()
		if err != nil {
			return nil, err
		}
		left = &ast.LogicalExpr{
			Token:    tok,
			Left:     left,
			Operator: token.AND,
			Right:    right,
		}
	}
	return left, nil
}
