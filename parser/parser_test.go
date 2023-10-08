package parser

import (
	"testing"

	"github.com/bbsemih/xgo/ast"
	"github.com/bbsemih/xgo/lexer"
)

func TestTanStatements(t *testing.T) {
	input := `
	tan a = 5;
	tan b = 10;
	tan xgo = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"xgo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testTanStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testTanStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "tan" {
		t.Errorf("s.TokenLiteral not 'tan'. got=%q", s.TokenLiteral())
		return false
	}

	tanStmt, ok := s.(*ast.TanStatement)
	if !ok {
		t.Errorf("s not *ast.TanStatement. got=%T", s)
		return false
	}

	if tanStmt.Name.Value != name {
		t.Errorf("tanStmt.Name.Value not '%s'. got=%s", name, tanStmt.Name.Value)
		return false
	}

	if tanStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, tanStmt.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 23;
	return 1283;
	return 999999;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program statements does not contain 3 statements. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

//45

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
