package lexer

import "github.com/azuyamat/hermit/internal/token"

func (l *Lexer) lexLiteral() (token.Token, error) {
	if tok, ok := l.lexFlag(); ok {
		return tok, nil
	}

	literal := l.readWord()
	if literal == "" {
		tok := token.New(token.ILLEGAL, string(l.ch), l.lineNumber, l.columnNumber)
		l.readChar()
		return tok, nil
	}

	tokenType := token.LookupIdent(literal)
	tok := token.New(tokenType, literal, l.lineNumber, l.columnNumber-len(literal))
	return tok, nil
}
