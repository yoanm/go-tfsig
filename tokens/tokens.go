package tokens

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// NewIdentTokens takes a string and convert it to hclwrite.Tokens containing a hclwrite.Token with hclsyntax.TokenIdent type
// See also NewIdentToken()
func NewIdentTokens(s string) hclwrite.Tokens {
	return hclwrite.Tokens{NewIdentToken([]byte(s))}
}

// NewCommaTokens creates a hclwrite.Tokens containing a hclwrite.Token with hclsyntax.TokenComma type
// See also NewCommaToken()
func NewCommaTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewCommaToken()}
}

// NewEqualTokens creates a hclwrite.Tokens containing a hclwrite.Token with hclsyntax.TokenEqual type
// See also NewEqualToken()
func NewEqualTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewEqualToken()}
}

// NewLineTokens creates a hclwrite.Tokens containing a hclwrite.Token with hclsyntax.TokenNewline type
// See also NewLineToken()
func NewLineTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewLineToken()}
}
