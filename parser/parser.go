package parser

import (
	"github.com/bbsemih/gomonk/token"

	"github.com/bbsemih/gomonk/lexer"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}
