package lexer

import (
	"testing"

	"github.com/azuyamat/hermit/internal/lexer"
	"github.com/azuyamat/hermit/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `ls -la | grep "test" > output.txt`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "ls"},
		{token.FLAG, "-la"},
		{token.PIPE, "|"},
		{token.IDENT, "grep"},
		{token.DOUBLE_QUOTED_STRING, "test"},
		{token.GT, ">"},
		{token.IDENT, "output"},
		{token.DOT, "."},
		{token.IDENT, "txt"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("test[%d] - unexpected error: %v", i, err)
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
