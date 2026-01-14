package lexer

import "github.com/azuyamat/hermit/internal/token"

func (l *Lexer) lexVariable() (token.Token, error) {
	var tok token.Token
	startCol := l.columnNumber
	if l.peekChar() == '(' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		tok = token.New(token.COMMAND_SUB_START, literal, l.lineNumber, startCol)
		l.readChar()
	} else if l.peekChar() == '{' {
		l.readChar() // consume $
		l.readChar() // consume {
		varName := l.readIdentifier()
		if l.ch != '}' {
			return token.Token{}, l.newError("unterminated variable expansion: expected }")
		}
		literal := "${" + varName + "}"
		tok = token.New(token.VARIABLE, literal, l.lineNumber, startCol)
	} else if isLetter(l.peekChar()) || l.peekChar() == '_' {
		l.readChar() // consume $
		varName := l.readIdentifier()
		literal := "$" + varName
		tok = token.New(token.VARIABLE, literal, l.lineNumber, startCol)
	} else {
		tok = token.New(token.ILLEGAL, string(l.ch), l.lineNumber, l.columnNumber)
	}
	return tok, nil
}
