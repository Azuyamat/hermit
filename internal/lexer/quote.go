package lexer

import "github.com/azuyamat/hermit/internal/token"

func (l *Lexer) lexQuotedString(quote byte) (token.Token, error) {
	var tok token.Token
	str, err := l.readString(quote)
	if err != nil {
		return token.Token{}, err
	}

	l.readChar()

	var tokenType token.TokenType
	switch quote {
	case '"':
		tokenType = token.DOUBLE_QUOTED_STRING
	case '\'':
		tokenType = token.SINGLE_QUOTED_STRING
	case '`':
		tokenType = token.BACKTICK_STRING
	}
	tok = token.New(tokenType, str, l.lineNumber, l.columnNumber)
	return tok, nil
}
