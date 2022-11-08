package tfsig

import (
	"regexp"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

var (
	invalidCharMatcher      *regexp.Regexp
	invalidFirstCharMatcher *regexp.Regexp
)

func init() {
	// From doc
	// > Identifiers can contain letters, digits, underscores (_), and hyphens (-).
	// > The first character of an identifier must not be a digit, to avoid ambiguity with literal numbers.
	// > For complete identifier rules, Terraform implements the Unicode identifier syntax, extended to include the ASCII hyphen character -.
	invalidCharMatcher = regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	invalidFirstCharMatcher = regexp.MustCompile(`^[0-9-]`)
}

// AppendBlockIfNotNil appends the provided block to the provided body only if block is not nil
// It simply avoids an if in your code
func AppendBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block) {
	if block != nil {
		body.AppendBlock(block)
	}
}

// AppendNewLineAndBlockIfNotNil appends an empty line followed by provided block to the provided body only if block is not nil
// It simply avoids an if in your code
func AppendNewLineAndBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block) {
	if block != nil {
		body.AppendNewline()
		body.AppendBlock(block)
	}
}

// ToTerraformIdentifier converts a string to a terraform identifier, by converting not allowed characters to '-'
// And if provided value starts with a character not allowed as first character, it replaces it by '_'
func ToTerraformIdentifier(s string) string {
	id := invalidCharMatcher.ReplaceAllString(s, "-")

	// Identifier must start with a letter or underscore !
	return invalidFirstCharMatcher.ReplaceAllString(id, "_")
}
