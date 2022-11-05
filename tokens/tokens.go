package tokens

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func NewIdentTokens(s string) hclwrite.Tokens {
	return hclwrite.Tokens{NewIdentToken([]byte(s))}
}

func NewLineTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewLineToken()}
}

func NewEqualTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewEqualToken()}
}

func NewCommaTokens() hclwrite.Tokens {
	return hclwrite.Tokens{NewCommaToken()}
}
