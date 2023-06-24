package token

import (
	"testing"
)

var keywordsTest = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func TestToken(t *testing.T) {
	if tok, ok := keywordsTest["let"]; ok {
		t.Logf("tok:%s, ok:%t", tok, ok)
	}
}
