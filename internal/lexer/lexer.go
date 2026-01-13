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
// Useful for lookahead operations such as distinguishing between '=' and '=='
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return EOF
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isIdentifierChar(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
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
	case '=':
		forwardIsEqual := l.peekChar() == '='
		if forwardIsEqual {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.EQ, literal, l.lineNumber, l.columnNumber-1)
		} else {
			tok = token.New(token.ASSIGN, string(l.ch), l.lineNumber, l.columnNumber)
		}
	case '+':
		tok = token.New(token.PLUS, string(l.ch), l.lineNumber, l.columnNumber)
	case '-':
		next := l.peekChar()
		if isLetter(next) || next == '-' {
			l.readChar()
			literal := "-"

			if next == '-' {
				literal += string(l.ch)
				l.readChar()
			}

			literal += l.readIdentifier()

			tok = token.New(token.FLAG, literal, l.lineNumber, l.columnNumber-len(literal))
			return tok, nil
		} else if isWhitespace(next) || next == EOF {
			tok = token.New(token.IDENT, "-", l.lineNumber, l.columnNumber)
		} else {
			tok = token.New(token.MINUS, string(l.ch), l.lineNumber, l.columnNumber)
		}
	case '!':
		forwardIsEqual := l.peekChar() == '='
		if forwardIsEqual {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.NOT_EQ, literal, l.lineNumber, l.columnNumber-1)
		} else {
			tok = token.New(token.BANG, string(l.ch), l.lineNumber, l.columnNumber)
		}
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch), l.lineNumber, l.columnNumber)
	case '/':
		tok = token.New(token.SLASH, string(l.ch), l.lineNumber, l.columnNumber)
	case '%':
		tok = token.New(token.PERCENT, string(l.ch), l.lineNumber, l.columnNumber)
	case '^':
		tok = token.New(token.CARET, string(l.ch), l.lineNumber, l.columnNumber)
	case '&':
		tok = token.New(token.AMPERSAND, string(l.ch), l.lineNumber, l.columnNumber)
	case '|':
		tok = token.New(token.PIPE, string(l.ch), l.lineNumber, l.columnNumber)
	case '~':
		tok = token.New(token.TILDE, string(l.ch), l.lineNumber, l.columnNumber)
	case '<':
		forwardIsLessThan := l.peekChar() == '<'
		forwardIsEqual := l.peekChar() == '='
		if forwardIsLessThan {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.LSHIFT, literal, l.lineNumber, l.columnNumber-1)
		} else if forwardIsEqual {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.LE, literal, l.lineNumber, l.columnNumber-1)
		} else {
			tok = token.New(token.LT, string(l.ch), l.lineNumber, l.columnNumber)
		}
	case '>':
		forwardIsGreaterThan := l.peekChar() == '>'
		forwardIsEqual := l.peekChar() == '='
		if forwardIsGreaterThan {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.RSHIFT, literal, l.lineNumber, l.columnNumber-1)
		} else if forwardIsEqual {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.GE, literal, l.lineNumber, l.columnNumber-1)
		} else {
			tok = token.New(token.GT, string(l.ch), l.lineNumber, l.columnNumber)
		}
	case ',':
		tok = token.New(token.COMMA, string(l.ch), l.lineNumber, l.columnNumber)
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch), l.lineNumber, l.columnNumber)
	case ':':
		tok = token.New(token.COLON, string(l.ch), l.lineNumber, l.columnNumber)
	case '?':
		tok = token.New(token.QUESTION, string(l.ch), l.lineNumber, l.columnNumber)
	case '.':
		tok = token.New(token.DOT, string(l.ch), l.lineNumber, l.columnNumber)
	case '(':
		tok = token.New(token.LPAREN, string(l.ch), l.lineNumber, l.columnNumber)
	case ')':
		tok = token.New(token.RPAREN, string(l.ch), l.lineNumber, l.columnNumber)
	case '{':
		tok = token.New(token.LBRACE, string(l.ch), l.lineNumber, l.columnNumber)
	case '}':
		tok = token.New(token.RBRACE, string(l.ch), l.lineNumber, l.columnNumber)
	case '[':
		tok = token.New(token.LBRACKET, string(l.ch), l.lineNumber, l.columnNumber)
	case ']':
		tok = token.New(token.RBRACKET, string(l.ch), l.lineNumber, l.columnNumber)
	case '"', '\'', '`':
		quote := l.ch
		str, err := l.readString(quote)
		if err != nil {
			return token.Token{}, err
		}
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
	case EOF:
		tok = token.New(token.EOF, "", l.lineNumber, l.columnNumber)
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := token.LookupIdent(literal)
			tok = token.New(tokenType, literal, l.lineNumber, l.columnNumber-len(literal))
			return tok, nil
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			tok = token.New(token.INT, literal, l.lineNumber, l.columnNumber-len(literal))
			return tok, nil
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), l.lineNumber, l.columnNumber)
		}
		return tok, nil
	}

	l.readChar()
	return tok, nil
}
