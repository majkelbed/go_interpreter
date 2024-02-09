package lexer

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
