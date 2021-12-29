package lexer 

import(
	"testing"
	"token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},`

	tests := []struct {
		expectedType  token.TokenType 
		expectedLiteral string 
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAR, "("},
		{token.RPAR, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `
	def add(x, y):
	    assert 0 <= x <= y
	    z = x + y
	    return z
	`

	tests := []struct {
		expectedType token.TokenType 
		expectedLiteral string 
	} {
		{token.DEF, "def"},
		{token.IDENTIFIER, "add"},
		{token.LPAR, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAR, ")"},
		{token.COLON, ":"},
		//第三节添加
		{token.ASSERT, "assert"},
		{token.NUMBER, "0"},
		{token.LESSEQUAL, "<="},
		{token.IDENTIFIER, "x"},
		{token.LESSEQUAL, "<="},
		{token.IDENTIFIER, "y"},
		//第三节添加
		{token.IDENTIFIER, "z"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "z"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "z"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `123abc`
	tests := []struct {
		expectedType token.TokenType 
		expectedLiteral string 
	} {
		{token.ILLEGAL, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {

			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}