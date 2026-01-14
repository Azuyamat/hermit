package lexer

import "github.com/azuyamat/hermit/internal/token"

func (l *Lexer) lexRedirect() (token.Token, error) {
	startLine := l.lineNumber
	startCol := l.columnNumber

	switch l.ch {
	case '>':
		if l.peekChar() == '>' {
			l.readChar()
			tok := token.New(token.REDIRECT_APPEND, ">>", startLine, startCol)
			l.readChar()
			return tok, nil
		} else if l.peekChar() == '&' {
			l.readChar()
			tok := token.New(token.REDIRECT_BOTH, ">&", startLine, startCol)
			l.readChar()
			return tok, nil
		}
		tok := token.New(token.REDIRECT_OUT, ">", startLine, startCol)
		l.readChar()
		return tok, nil

	case '<':
		tok := token.New(token.REDIRECT_IN, "<", startLine, startCol)
		l.readChar()
		return tok, nil

	case '&':
		l.readChar() // consume '>'
		tok := token.New(token.REDIRECT_BOTH, "&>", startLine, startCol)
		l.readChar()
		return tok, nil

	case '2':
		l.readChar() // consume '>'
		if l.peekChar() == '>' {
			l.readChar() // consume second '>'
			tok := token.New(token.REDIRECT_ERR_APPEND, "2>>", startLine, startCol)
			l.readChar()
			return tok, nil
		}
		tok := token.New(token.REDIRECT_ERR, "2>", startLine, startCol)
		l.readChar()
		return tok, nil
	}

	return token.New(token.ILLEGAL, string(l.ch), startLine, startCol), nil
}
