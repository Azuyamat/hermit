package lexer

import "fmt"

type Error struct {
	Pos     Position
	Message string
}

type Position struct {
	Line   int
	Column int
	Offset int
}

func (l *Lexer) currentPosition() Position {
	return Position{
		Line:   l.lineNumber,
		Column: l.columnNumber,
		Offset: l.position,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("lexer error at line %d, column %d: %s", e.Pos.Line, e.Pos.Column, e.Message)
}

type ErrUnterminatedString struct {
	Pos   Position
	Quote byte
}

func (e *ErrUnterminatedString) Error() string {
	return fmt.Sprintf("unterminated string starting at line %d, column %d with quote %q", e.Pos.Line, e.Pos.Column, e.Quote)
}

func (l *Lexer) newError(message string) *Error {
	return &Error{
		Pos:     l.currentPosition(),
		Message: message,
	}
}

func (l *Lexer) newUnterminatedStringError(quote byte) *ErrUnterminatedString {
	return &ErrUnterminatedString{
		Pos:   l.currentPosition(),
		Quote: quote,
	}
}
