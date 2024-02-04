package parser

import (
	"github.com/makarellav/monkey-go/ast"
	"github.com/makarellav/monkey-go/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("returnStmt not *ast.ReturnStatement. got=%T", returnStmt)

			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got=%q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.Errors()) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(p.Errors()))

	for _, err := range p.Errors() {
		t.Errorf("parser error: %q", err)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())

		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", letStmt)

		return false
	}

	if letStmt.Name.Value != name {
		t.Fatalf("letStmt.Name.Value not %q. got=%q", name, letStmt.Name.Value)

		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("letStmt.Name.TokenLiteral() not %q. got=%q", name, letStmt.Name.TokenLiteral())

		return false
	}

	return true
}
