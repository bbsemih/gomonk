package parser

import (
	"github.com/bbsemih/gomonk/ast"

	"github.com/bbsemih/gomonk/token"

	"github.com/bbsemih/gomonk/lexer"
)

type Parser struct {
	// l is a pointer to an instance of the lexer
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
