package tfsig

import (
	"github.com/zclconf/go-cty/cty"
)

/** Public **/

// AppendAttrIfNotNil appends the provided attribute only if not nil
//
// It simply avoids an `if` in your code.
func (s *BlockSignature) AppendAttrIfNotNil(attrName string, v *cty.Value) {
	if v != nil {
		s.AppendAttribute(attrName, *v)
	}
}

// AppendChildIfNotNil appends the provided child only if not nil. And in case there is existing elements,
// it prepends an empty line
//
// It simply avoids an `if` in your code.
func (s *BlockSignature) AppendChildIfNotNil(child *BlockSignature) {
	if child != nil {
		if len(s.GetElements()) > 0 {
			s.AppendEmptyLine()
		}

		s.AppendChild(child)
	}
}
