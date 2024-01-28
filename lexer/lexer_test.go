package lexer

import (
	"app/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let let_int = 123;

		const const_lambda = () => {
			if (let_int < 100) {
				return true;
			} else if (let_int == 100) {
				return false;
			} else {
				return null;
			}
		};
		
		var var_fn = fn () {
			return "some string";
		};
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "let_int"},
		{token.ASSIGN, "="},
		{token.INT, "123"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "const_lambda"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LAMBDA, "=>"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "let_int"},
		{token.LT, "<"},
		{token.INT, "100"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE_IF, "else if"},
		{token.LPAREN, "("},
		{token.IDENT, "let_int"},
		{token.EQ, "=="},
		{token.INT, "100"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.NULL, "null"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.VAR, "var"},
		{token.IDENT, "var_fn"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.STRING, "some string"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(string(input))

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
