package ast

// abstract syntax tree is a tree representation of the abstract syntactic structure of source code
// this needed to be done because we need to parse the code and create a tree representation of it
// source code > lexer > tokens > parser > AST > interpreter

import "github.com/bbsemih/xgo/token"

// every node in AST has to implement the Node interface
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// every valid xgo program is a series of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type TanStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ls *TanStatement) statementNode() {}
func (ls *TanStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
