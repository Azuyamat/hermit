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
	FLAG  = "FLAG"

	DOUBLE_QUOTED_STRING = "DOUBLE_QUOTED_STRING"
	SINGLE_QUOTED_STRING = "SINGLE_QUOTED_STRING"
	BACKTICK_STRING      = "BACKTICK_STRING"

	PIPE = "|"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	AND = "AND" // &&
	OR  = "OR"  // ||

	COMMAND_SUB_START = "$("

	REDIRECT_OUT        = "REDIRECT_OUT"        // >
	REDIRECT_APPEND     = "REDIRECT_APPEND"     // >>
	REDIRECT_IN         = "REDIRECT_IN"         // <
	REDIRECT_ERR        = "REDIRECT_ERR"        // 2>
	REDIRECT_ERR_APPEND = "REDIRECT_ERR_APPEND" // 2>>
	REDIRECT_BOTH       = "REDIRECT_BOTH"       // &>

	VARIABLE = "VARIABLE"
)

var keywords = map[string]TokenType{
	"if":       "IF",
	"then":     "THEN",
	"else":     "ELSE",
	"elif":     "ELIF",
	"fi":       "FI",
	"for":      "FOR",
	"while":    "WHILE",
	"do":       "DO",
	"done":     "DONE",
	"case":     "CASE",
	"esac":     "ESAC",
	"function": "FUNCTION",
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func New(tokenType TokenType, literal string, lineNumber, columnNumber int) Token {
	return Token{
		Type:         tokenType,
		Literal:      literal,
		LineNumber:   lineNumber,
		ColumnNumber: columnNumber,
	}
}
