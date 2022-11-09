package tfsig

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/tokens"
)

/** Public **/

// DependsOn adds an empty line and the 'depends_on' terraform directive with provided id list
func (s *BlockSignature) DependsOn(idList []string) {
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

func (c *LifecycleConfig) SetCreateBeforeDestroy(b bool) {
	c.CreateBeforeDestroy = &b
}
func (c *LifecycleConfig) SetPreventDestroy(b bool) {
	c.PreventDestroy = &b
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
	cond.AppendAttribute("error_message", cty.StringVal(c.errorMessage))

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
