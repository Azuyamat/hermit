package token

type TokenType string

type Token struct {
	Type         TokenType
	Literal      string
	LineNumber   int
	ColumnNumber int
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"
	INT   = "INT"
	FLAG  = "FLAG"

	DOUBLE_QUOTED_STRING = "DOUBLE_QUOTED_STRING"
	SINGLE_QUOTED_STRING = "SINGLE_QUOTED_STRING"
	BACKTICK_STRING      = "BACKTICK_STRING"

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	PERCENT   = "%"
	CARET     = "^"
	AMPERSAND = "&"
	PIPE      = "|"
	TILDE     = "~"
	LSHIFT    = "<<"
	RSHIFT    = ">>"

	LT = "<"
	GT = ">"
	LE = "<="
	GE = ">="

	EQ     = "=="
	NOT_EQ = "!="

	QUESTION  = "?"
	COLON     = ":"
	COMMA     = ","
	SEMICOLON = ";"

	DOT = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
)

func New(tokenType TokenType, literal string, lineNumber, columnNumber int) Token {
	return Token{
		Type:         tokenType,
		Literal:      literal,
		LineNumber:   lineNumber,
		ColumnNumber: columnNumber,
	}
}
