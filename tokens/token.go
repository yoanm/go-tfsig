package tokens

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func NewIdentToken(b []byte) *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: b}
}

func NewCommaToken() *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte{','}}
}

func NewEqualToken() *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte("=")}
}

func NewLineToken() *hclwrite.Token {
	return &hclwrite.Token{
		Type:  hclsyntax.TokenNewline,
		Bytes: []byte{'\n'},
	}
}
