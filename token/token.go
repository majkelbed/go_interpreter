package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	PLUS_EQ  = "+="
	MINUS    = "-"
	MINUS_EQ = "-="
	EQ       = "=="
	LT       = "<"
	LT_EQ    = "<="
	GT       = ">"
	GT_EQ    = ">="
	LAMBDA   = "=>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	VAR      = "VAR"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE_IF  = "ELSE_IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"let":     LET,
	"const":   CONST,
	"var":     VAR,
	"return":  RETURN,
	"if":      IF,
	"else if": ELSE_IF,
	"else":    ELSE,
	"true":    TRUE,
	"false":   FALSE,
	"null":    NULL,
}

func LookupIdent(ident string) TokenType {
	// TODO: this code is strange, check docs for this weird IF
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func New(tokenType TokenType, literal interface{}) Token {
	switch v := literal.(type) {
	case string:
		return Token{Type: tokenType, Literal: v}
	case byte:
		return Token{Type: tokenType, Literal: string(v)}
	default:
		panic(v)
	}
}
