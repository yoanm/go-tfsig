/*
Package tfsig is a wrapper for Terraform HCL language (hclwrite).
It provides ability to generate block signature which are way easier to manipulate and alter than hclwrite.tokens type
*/
package tfsig

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/tokens"
)

/** Public **/

// NewSignature returns a BlockSignature pointer filled with provided labels and elements
func NewSignature(name string, labels []string, elements BodyElements) *BlockSignature {
	return &BlockSignature{
		typeName: name,
		labels:   labels,
		elements: elements,
	}
}

// NewEmptySignature returns a BlockSignature pointer filled with provided labels
func NewEmptySignature(name string, labels ...string) *BlockSignature {
	return NewSignature(name, labels, BodyElements{})
}

// NewEmptyResource returns a BlockSignature pointer with "resource" type and filled with provided labels
func NewEmptyResource(name, id string, labels ...string) *BlockSignature {
	return NewEmptySignature("resource", append([]string{name, id}, labels...)...)
}

// BlockSignature is basically a wrapper to HCL blocks
// It holds a type, the block labels and its elements
type BlockSignature struct {
	typeName string
	labels   []string
	elements BodyElements
}

// GetType returns the type of the block
func (s *BlockSignature) GetType() string {
	return s.typeName
}

// GetLabels returns labels attached to the block
func (s *BlockSignature) GetLabels() []string {
	return s.labels
}

// GetElements returns all elements attached to the block
func (s *BlockSignature) GetElements() BodyElements {
	return s.elements
}

// SetElements overrides existing elements by provided ones
func (s *BlockSignature) SetElements(elements BodyElements) {
	s.elements = elements
}

// AppendElement appends an element to the block
func (s *BlockSignature) AppendElement(element BodyElement) {
	s.elements = append(s.elements, element)
}

// AppendAttribute appends an attribute to the block
func (s *BlockSignature) AppendAttribute(name string, value cty.Value) {
	s.AppendElement(NewBodyAttribute(name, value))
}

// AppendChild appends a child block to the block
func (s *BlockSignature) AppendChild(child *BlockSignature) {
	s.AppendElement(NewBodyBlock(child))
}

// AppendEmptyLine appends an empty line to the block
func (s *BlockSignature) AppendEmptyLine() {
	s.AppendElement(NewBodyEmptyLine())
}

// Build creates a hclwrite.Block and appends block's elements to it
func (s *BlockSignature) Build() *hclwrite.Block {
	block := hclwrite.NewBlock(s.GetType(), s.GetLabels())

	s.writeElementsToBody(block.Body())

	return block
}

// BuildTokens builds the block signature as hclwrite.Tokens
func (s *BlockSignature) BuildTokens() (tks hclwrite.Tokens) {
	if block := s.Build(); block != nil {
		blockTks := block.BuildTokens(nil)
		// Remove trailing new line automatically added (=remove last token)
		tks = append(hclwrite.Tokens{}, blockTks[0:len(blockTks)-1]...)
	}

	return tks
}

/** Private **/

// writeElementsToBody writes all block signature elements to the provided hclwrite.Body
//
// It takes care of attribute values containing hclwrite.Tokens encapsulated into a cty capsule
func (s *BlockSignature) writeElementsToBody(body *hclwrite.Body) {
	for _, value := range s.GetElements() {
		if value.IsBodyBlock() {
			body.AppendBlock(value.Build())
		} else if value.IsBodyAttribute() {
			if tokens.ContainsCapsule(value.attr) {
				body.SetAttributeRaw(value.GetName(), tokens.Generate(value.attr))
			} else {
				body.SetAttributeValue(value.GetName(), *value.attr)
			}
		} else if value.IsBodyEmptyLine() {
			body.AppendNewline()
		}
	}
}
