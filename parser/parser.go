package parser

import (
	"app/ast"
	"app/lexer"
	"app/token"
	"errors"
	"fmt"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	nextToken    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.advance()
	p.advance()

	return p
}

func (p *Parser) advance() {
	p.currentToken = p.nextToken
	p.nextToken = p.l.Advance()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		// TODO: add assigning new value to existing variable
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectNext(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.nextTokenIs(token.SEMICOLON) {
		p.advance()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	for !p.nextTokenIs(token.SEMICOLON) {
		p.advance()
	}

	return stmt
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}

func (p *Parser) addTokenTypeMismatchError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expected %s got %s instead", t, p.nextToken.Type))
}

func (p *Parser) expectNext(t token.TokenType) bool {
	if p.nextTokenIs(t) {
		p.advance()
		return true
	} else {
		p.addTokenTypeMismatchError(t)
		return false
	}
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.advance()
	}

	if len(p.errors) > 0 {
		var errs []error

		for _, e := range p.errors {
			err := errors.New(e)
			errs = append(errs, err)
		}

		return nil, errors.Join(errs...)
	}

	return program, nil
}
