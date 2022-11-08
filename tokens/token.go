package tokens

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// NewIdentToken returns a hclwrite.Token with hclsyntax.TokenIdent type encapsulating provided bytes
func NewIdentToken(b []byte) *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: b}
}

// NewCommaToken returns a hclwrite.Token with hclsyntax.TokenComma type
func NewCommaToken() *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte{','}}
}

// NewEqualToken returns a hclwrite.Token with hclsyntax.TokenEqual type
func NewEqualToken() *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte("=")}
}

// NewLineToken returns a hclwrite.Token with hclsyntax.TokenNewline type
func NewLineToken() *hclwrite.Token {
	return &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}}
}
