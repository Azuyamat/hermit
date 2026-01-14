package lexer

import "github.com/azuyamat/hermit/internal/token"

func (l *Lexer) lexFlag() (token.Token, bool) {
	if l.ch != '-' {
		return token.Token{}, false
	}

	next := l.peekChar()
	if isLetter(next) || next == '-' {
		literal := l.readWord()
		tok := token.New(token.FLAG, literal, l.lineNumber, l.columnNumber-len(literal))
		return tok, true
	}

	return token.Token{}, false
}
