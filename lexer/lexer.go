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

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) Advance() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		next := l.peekChar()
		if next == '>' {
			t = token.New(token.LAMBDA, string(l.ch)+string(next))
			l.readChar()
		} else if next == '=' {
			t = token.New(token.EQ, string(l.ch)+string(next))
			l.readChar()
		} else {
			t = token.New(token.ASSIGN, l.ch)
		}
	case ';':
		t = token.New(token.SEMICOLON, l.ch)
	case '(':
		t = token.New(token.LPAREN, l.ch)
	case ')':
		t = token.New(token.RPAREN, l.ch)
	case ',':
		t = token.New(token.COMMA, l.ch)
	case '+':
		if l.peekChar() == '=' {
			t = token.New(token.PLUS_EQ, l.ch)
			l.readChar()
		} else {
			t = token.New(token.PLUS, l.ch)
		}
	case '{':
		t = token.New(token.LBRACE, l.ch)
	case '}':
		t = token.New(token.RBRACE, l.ch)
	case '<':
		t = token.New(token.LT, l.ch)
	case '>':
		t = token.New(token.GT, l.ch)
	case '"':
		l.readChar()
		str := ""

		for l.ch != '"' {
			str += string(l.ch)
			l.readChar()
		}

		t = token.New(token.STRING, str)
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
			t = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return t
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}
