package tfsig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func NewBodyBlock(block *BlockSignature) BodyElement {
	return BodyElement{name: block.GetType(), block: block, attr: nil, isEmptyLine: false}
}
func NewBodyAttribute(name string, attr cty.Value) BodyElement {
	return BodyElement{name: name, attr: &attr, isEmptyLine: false}
}
func NewBodyEmptyLine() BodyElement {
	return BodyElement{name: "empty_line", block: nil, attr: nil, isEmptyLine: true}
}

type BodyElement struct {
	name        string
	block       *BlockSignature
	attr        *cty.Value
	isEmptyLine bool
}

type BodyElements []BodyElement

func (e BodyElement) GetName() string {
	return e.name
}
func (e BodyElement) IsBodyBlock() bool {
	return e.block != nil
}

func (e BodyElement) IsBodyAttribute() bool {
	return e.attr != nil
}
func (e BodyElement) IsBodyEmptyLine() bool {
	return e.isEmptyLine
}

// GetBodyAttribute will panic if BodyElement is not an attribute
func (e BodyElement) GetBodyAttribute() *cty.Value {
	if !e.IsBodyAttribute() {
		panic("not a body attribute")
	}
	return e.attr
}

// GetBodyBlock will panic if BodyElement is not a block
func (e BodyElement) GetBodyBlock() *BlockSignature {
	if !e.IsBodyBlock() {
		panic("not a body block")
	}
	return e.block
}

// Build will panic if element is not a block !
func (e BodyElement) Build() *hclwrite.Block {
	if !e.IsBodyBlock() {
		panic(fmt.Sprintf("Element is not a block"))
	}

	return e.block.Build()
}
