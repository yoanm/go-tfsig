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

	s.WriteElementsToBody(block.Body())

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

// WriteElementsToBody writes all block signature elements to the provided hclwrite.Body
//
// It takes care of attribute values containing hclwrite.Tokens encapsulated into a cty capsule
func (s *BlockSignature) WriteElementsToBody(body *hclwrite.Body) {
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

// DependsOn adds an empty line and the 'depends_on' terraform directive with provided id list
func (s *BlockSignature) DependsOn(idList []string) {
	if idList == nil {
		return
	}

	s.AppendEmptyLine()
	s.AppendAttribute("depends_on", *tokens.NewIdentListValue(idList))
}

// LifecycleConfig is used as argument for Lifecycle() method
// It's basically a wrapper for terraform 'lifecycle' directive
type LifecycleConfig struct {
	CreateBeforeDestroy *bool
	PreventDestroy      *bool
	IgnoreChanges       []string
	ReplaceTriggeredBy  []string
	Precondition        *LifecycleCondition
	Postcondition       *LifecycleCondition
}

// LifecycleCondition is used for Precondition and Postcondition property of LifecycleConfig
// It's basically a wrapper for terraform lifecycle pre- and post-conditions
type LifecycleCondition struct {
	condition    string
	errorMessage string
}

// Lifecycle adds an empty line and the 'lifecycle' terraform directive and then append provided lifecycle attributes
func (s *BlockSignature) Lifecycle(config LifecycleConfig) {
	sig := NewEmptySignature("lifecycle")

	appendLifecycleBoolAttribute(sig, "create_before_destroy", config.CreateBeforeDestroy)
	appendLifecycleBoolAttribute(sig, "prevent_destroy", config.PreventDestroy)

	if config.IgnoreChanges != nil {
		sig.AppendAttribute("ignore_changes", *tokens.NewIdentListValue(config.IgnoreChanges))
	}

	appendLifecycleConditionBlock(sig, "precondition", config.Precondition)
	appendLifecycleConditionBlock(sig, "postcondition", config.Postcondition)

	s.AppendEmptyLine()
	s.AppendChild(sig)
}

/** Private **/

func appendLifecycleConditionBlock(lifecycleSig *BlockSignature, name string, c *LifecycleCondition) {
	if c == nil {
		return
	}
	cond := NewEmptySignature(name)

	cond.AppendAttribute("condition", *tokens.NewIdentValue(c.condition))
	cond.AppendAttribute("error_message", *tokens.NewIdentValue(c.errorMessage))

	lifecycleSig.AppendChild(cond)
}

func appendLifecycleBoolAttribute(lifecycleSig *BlockSignature, name string, value *bool) {
	if value == nil {
		return
	}
	val := "false"
	if *value {
		val = "true"
	}
	lifecycleSig.AppendAttribute(name, *tokens.NewIdentValue(val))
}
