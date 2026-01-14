package lexer

import (
	"github.com/azuyamat/hermit/internal/token"
)

const EOF = 0

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
	lineNumber   int
	columnNumber int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, lineNumber: 1, columnNumber: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.columnNumber++
}

// See the next character without moving the position forward
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return EOF
	} else {
		return l.input[l.readPosition]
	}
}

// readWord reads a continuous sequence of non-whitespace, non-operator characters
// This handles paths like README.md, ../file.txt, /usr/bin/cat as single tokens
func (l *Lexer) readWord() string {
	start := l.position
	for !l.isDelimiter(l.ch) && l.ch != EOF {
		l.readChar()
	}
	return l.input[start:l.position]
}

// readIdentifier reads a valid identifier (letters, digits, underscores)
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.position]
}

// isDelimiter returns true for characters that separate tokens
func (l *Lexer) isDelimiter(ch byte) bool {
	return isWhitespace(ch) ||
		ch == '|' || ch == '>' || ch == '<' ||
		ch == '&' || ch == ';' ||
		ch == '(' || ch == ')' ||
		ch == '{' || ch == '}' ||
		ch == '[' || ch == ']' ||
		ch == '"' || ch == '\'' || ch == '`' ||
		ch == '$' // Stop at $ for command substitution
}

func (l *Lexer) readString(quote byte) (string, error) {
	start := l.position + 1 // skip the opening quote

	for {
		l.readChar()

		if l.ch == EOF {
			return "", l.newUnterminatedStringError(quote)
		}

		if l.ch == '\\' {
			l.readChar() // skip the escaped character
			continue
		}

		if l.ch == quote {
			break
		}
	}

	return l.input[start:l.position], nil
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		if isNewline(l.ch) {
			l.lineNumber++
			l.columnNumber = 0
		}
		l.readChar()
	}
}

func (l *Lexer) NextToken() (token.Token, error) {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.OR, string(ch)+string(l.ch), l.lineNumber, l.columnNumber)
		} else {
			tok = token.New(token.PIPE, string(l.ch), l.lineNumber, l.columnNumber)
		}
		l.readChar()
	case '(':
		tok = token.New(token.LPAREN, string(l.ch), l.lineNumber, l.columnNumber)
		l.readChar()
	case ')':
		tok = token.New(token.RPAREN, string(l.ch), l.lineNumber, l.columnNumber)
		l.readChar()
	case '{':
		tok = token.New(token.LBRACE, string(l.ch), l.lineNumber, l.columnNumber)
		l.readChar()
	case '}':
		tok = token.New(token.RBRACE, string(l.ch), l.lineNumber, l.columnNumber)
		l.readChar()
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.AND, string(ch)+string(l.ch), l.lineNumber, l.columnNumber)
		} else if l.peekChar() == '>' {
			return l.lexRedirect()
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), l.lineNumber, l.columnNumber)
		}
		l.readChar()
	case '$':
		return l.lexVariable()
	case '"', '\'', '`':
		return l.lexQuotedString(l.ch)
	case '>', '<':
		return l.lexRedirect()
	case '2':
		if l.peekChar() == '>' {
			return l.lexRedirect()
		}
		return l.lexLiteral()
	case EOF:
		tok = token.New(token.EOF, "", l.lineNumber, l.columnNumber)
		return tok, nil
	default:
		return l.lexLiteral()
	}

	return tok, nil
}
