package tfsig

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// AppendBlockIfNotNil appends the provided block to the provided body only if block is not nil
//
// It simply avoids an `if` in your code.
func AppendBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block) {
	if block != nil {
		body.AppendBlock(block)
	}
}

// AppendNewLineAndBlockIfNotNil appends an empty line followed by provided block to the provided body
// only if block is not nil
//
// It simply avoids an `if` in your code.
func AppendNewLineAndBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block) {
	if block != nil {
		body.AppendNewline()
		body.AppendBlock(block)
	}
}

// AppendAttributeIfNotNil appends the provided attribute to the signature only if not nil
//
// It simply avoids an `if` in your code.
func AppendAttributeIfNotNil(sig *BlockSignature, attrName string, v *cty.Value) {
	if v != nil {
		sig.AppendAttribute(attrName, *v)
	}
}

// AppendChildIfNotNil appends the provided child to the signature only if not nil.
// And in case there is existing elements, it prepends an empty line
//
// It simply avoids two `if` in your code.
func AppendChildIfNotNil(sig *BlockSignature, child *BlockSignature) {
	if child != nil {
		if len(sig.GetElements()) > 0 {
			sig.AppendEmptyLine()
		}

		sig.AppendChild(child)
	}
}
