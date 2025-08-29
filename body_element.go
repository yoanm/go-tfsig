package tfsig

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// NewBodyBlock returns a Block BodyElement.
func NewBodyBlock(block *BlockSignature) BodyElement {
	return BodyElement{name: block.GetType(), block: block, isEmptyLine: false, attr: nil}
}

// NewBodyAttribute returns an Attribute BodyElement.
func NewBodyAttribute(name string, attr cty.Value) BodyElement {
	return BodyElement{name: name, attr: &attr, isEmptyLine: false, block: nil}
}

// NewBodyEmptyLine returns an empty line BodyElement.
func NewBodyEmptyLine() BodyElement {
	return BodyElement{name: "empty_line", isEmptyLine: true, block: nil, attr: nil}
}

// BodyElement is a wrapper for more or less anything that can be appended to a BlockSignature.
type BodyElement struct {
	name        string
	block       *BlockSignature
	attr        *cty.Value
	isEmptyLine bool
}

// BodyElements is a simple wrapper for a list of BodyElement.
type BodyElements []BodyElement

// GetName returns the name of the BodyElement.
func (e BodyElement) GetName() string {
	return e.name
}

// IsBodyBlock returns true if the BodyElement is a block.
func (e BodyElement) IsBodyBlock() bool {
	return e.block != nil
}

// IsBodyAttribute returns true if the BodyElement is an attribute.
func (e BodyElement) IsBodyAttribute() bool {
	return e.attr != nil
}

// IsBodyEmptyLine returns true if the BodyElement is an empty line.
func (e BodyElement) IsBodyEmptyLine() bool {
	return e.isEmptyLine
}

// GetBodyAttribute returns the value of the attribute behind the BodyElement
//
// It panics if BodyElement is not an attribute (use `IsBodyAttribute()` first).
func (e BodyElement) GetBodyAttribute() *cty.Value {
	if !e.IsBodyAttribute() {
		panic("element is not a body attribute")
	}

	return e.attr
}

// GetBodyBlock returns the block behind the BodyElement
//
// it panics if BodyElement is not a block (use `IsBodyBlock()` first).
func (e BodyElement) GetBodyBlock() *BlockSignature {
	if !e.IsBodyBlock() {
		panic("element is not a body block")
	}

	return e.block
}

func (e BodyElement) GetBodyBlock2() *BlockSignature {
	if !e.IsBodyBlock() {
		panic("element is not a body block")
	}

	return e.block
}

// Build convert the current BodyElement into a `hclwrite.Block`
//
// it panics if BodyElement is not a block (use `IsBodyBlock()` first).
func (e BodyElement) Build() *hclwrite.Block {
	if !e.IsBodyBlock() {
		panic("element is not a body block")
	}

	return e.block.Build()
}
