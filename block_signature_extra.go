package tfsig

import (
	"github.com/zclconf/go-cty/cty"
)

/** Public **/

// AppendAttrIfNotNil appends the provided attribute only if not nil
//
// It simply avoids an `if` in your code.
func (sig *BlockSignature) AppendAttributeIfNotNil(attrName string, v *cty.Value) {
	if v != nil {
		sig.AppendAttribute(attrName, *v)
	}
}

// AppendChildIfNotNil appends the provided child only if not nil. And in case there is existing elements,
// it prepends an empty line
//
// It simply avoids an `if` in your code.
func (sig *BlockSignature) AppendChildIfNotNil(child *BlockSignature) {
	if child != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendChild(child)
	}
}
