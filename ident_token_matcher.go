package tfsig

import (
	"strings"
)

//nolint:gochecknoglobals // Better to keep it as **internal** global var than define it each time
var defaultUnescapedStringPrefixList = []string{
	// Common terraform prefix for unescaped strings - START
	"local.",
	"var.",
	"data.",
	// Common terraform prefix for unescaped strings - END
}

// NewIdentTokenMatcher returns an instance of IdentTokenMatcher with provided list of prefix to
// consider as 'ident' tokens
//
// `local.`, `var.` and `data.` tokens will be considered as 'ident' tokens by default.
func NewIdentTokenMatcher(prefixList ...string) IdentTokenMatcher {
	return IdentTokenMatcher{prefixList: append(prefixList, defaultUnescapedStringPrefixList...)}
}

// IdentTokenMatcherInterface is a simple interface declaring required method to detect an 'ident' token.
type IdentTokenMatcherInterface interface {
	IsIdentToken(s string) bool
}

// IdentTokenMatcher is a simple implementation for IdentTokenMatcherInterface.
type IdentTokenMatcher struct {
	prefixList []string
}

// IsIdentToken is the implementation for IdentTokenMatcherInterface.
func (m IdentTokenMatcher) IsIdentToken(s string) bool {
	for _, prefix := range m.prefixList {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}
