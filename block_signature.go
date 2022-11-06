package tfsig

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/tokens"
)

/** Public **/

func NewSignature(name string, labels []string, elements BodyElements) *BlockSignature {
	return &BlockSignature{
		typeName: name,
		labels:   labels,
		elements: elements,
	}
}
func NewEmptySignature(name string, labels ...string) *BlockSignature {
	return NewSignature(name, labels, BodyElements{})
}

func NewEmptyResource(name, id string, labels ...string) *BlockSignature {
	return NewEmptySignature("resource", append([]string{name, id}, labels...)...)
}

type BlockSignature struct {
	typeName string
	labels   []string
	elements BodyElements
}

func (signature *BlockSignature) GetType() string {
	return signature.typeName
}
func (signature *BlockSignature) GetLabels() []string {
	return signature.labels
}
func (signature *BlockSignature) GetElements() BodyElements {
	return signature.elements
}
func (signature *BlockSignature) SetElements(elements BodyElements) {
	signature.elements = elements
}
func (signature *BlockSignature) AppendElement(element BodyElement) {
	signature.elements = append(signature.elements, element)
}

func (signature *BlockSignature) AppendAttribute(name string, value cty.Value) {
	signature.AppendElement(NewBodyAttribute(name, value))
}
func (signature *BlockSignature) AppendChild(child *BlockSignature) {
	signature.AppendElement(NewBodyBlock(child))
}
func (signature *BlockSignature) AppendEmptyLine() {
	signature.AppendElement(NewBodyEmptyLine())
}

func (signature *BlockSignature) Build() *hclwrite.Block {
	block := hclwrite.NewBlock(signature.GetType(), signature.GetLabels())

	signature.WriteElementsToBody(block.Body())

	return block
}

func (signature *BlockSignature) BuildTokens() (tks hclwrite.Tokens) {
	if block := signature.Build(); block != nil {
		blockTks := block.BuildTokens(nil)
		// Remove trailing new line automatically added (=remove last token)
		tks = append(hclwrite.Tokens{}, blockTks[0:len(blockTks)-1]...)
	}

	return tks
}

func (signature *BlockSignature) WriteElementsToBody(body *hclwrite.Body) {
	for _, value := range signature.GetElements() {
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

// DependsOn adds an empty line and the 'depends_on terraform directive with provided id list
func (s *BlockSignature) DependsOn(idList []string) {
	if idList == nil {
		return
	}

	s.AppendEmptyLine()
	s.AppendAttribute("depends_on", *tokens.NewIdentListValue(idList))
}

type LifecycleCondition struct {
	condition    string
	errorMessage string
}
type LifecycleConfig struct {
	CreateBeforeDestroy *bool
	PreventDestroy      *bool
	IgnoreChanges       []string
	ReplaceTriggeredBy  []string
	Precondition        *LifecycleCondition
	Postcondition       *LifecycleCondition
}

// Lifecycle adds an empty line and the 'lifecycle terraform directive, or return the existing signature
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
