// lexer/lexer.go
package lexer

import (
	"app/token"
	"regexp"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

var validDigit = regexp.MustCompile(`[0-9]`)

// TODO: replace all this mambo jambo with a simple regexp lexer

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	// Why not add white space as case and just skip it this way?
	l.skipWhitespace()

	switch l.ch {
	case '=':
		next := l.peekChar()
		if next == '>' {
			t = newToken(token.LAMBDA, string(l.ch)+string(next))
			l.readChar()
		} else if next == '=' {
			t = newToken(token.EQ, string(l.ch)+string(next))
			l.readChar()
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		if l.peekChar() == '=' {
			t = newToken(token.PLUS_EQ, l.ch)
			l.readChar()
		} else {
			t = newToken(token.PLUS, l.ch)
		}
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case '"':
		l.readChar()
		str := ""

		for l.ch != '"' {
			str += string(l.ch)
			l.readChar()
		}

		t = newToken(token.STRING, str)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()

			if ident == "else" {
				l.skipWhitespace()

				if isLetter(l.ch) {
					nextIdent := l.readIdentifier()

					if nextIdent == "if" {
						ident = "else if"
					}
				}
			}

			t.Literal = ident
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return t
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, literal interface{}) token.Token {
	switch v := literal.(type) {
	case string:
		return token.Token{Type: tokenType, Literal: v}
	case byte:
		return token.Token{Type: tokenType, Literal: string(v)}
	default:
		panic(v)
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}
