package tfsig

import (
	"strings"
)

var defaultUnescapedStringPrefixList []string

func init() {
	defaultUnescapedStringPrefixList = []string{
		// Common terraform prefix for unescaped strings - START
		"local.",
		"var.",
		"data.",
		// Common terraform prefix for unescaped strings - END
	}
}

type IdentTokenMatcher struct {
	prefixList []string
}

func NewIdentTokenMatcher(prefixList ...string) IdentTokenMatcher {
	return IdentTokenMatcher{prefixList: append(prefixList, defaultUnescapedStringPrefixList...)}
}

func (m *IdentTokenMatcher) IsIdentToken(s string) bool {
	for _, prefix := range m.prefixList {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}
