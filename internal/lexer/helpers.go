package lexer

const SEPARATOR = '_'

func isIdentifierChar(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || isSeparator(ch)
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isNewline(ch byte) bool {
	return ch == '\n' || ch == '\r'
}

func isSeparator(ch byte) bool {
	return ch == SEPARATOR
}
