package tfsig

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/tokens"
)

/** Public **/

// DependsOn adds an empty line and the 'depends_on' terraform directive with provided id list.
func (sig *BlockSignature) DependsOn(idList []string) {
	sig.AppendEmptyLine()
	sig.AppendAttribute("depends_on", *tokens.NewIdentListValue(idList))
}

// LifecycleConfig is used as argument for `Lifecycle()` method
// It's basically a wrapper for terraform `lifecycle` directive.
type LifecycleConfig struct {
	CreateBeforeDestroy *bool
	PreventDestroy      *bool
	IgnoreChanges       []string
	ReplaceTriggeredBy  []string
	Precondition        *LifecycleCondition
	Postcondition       *LifecycleCondition
}

// SetCreateBeforeDestroy is a simple helper to avoid having to create a boolean variable
// and then pass the pointer to it
//
// E.g: instead of writing
//
//	createBeforeDestroy = true
//	config := LifecycleConfig{CreateBeforeDestroy: &createBeforeDestroy}
//
// Simply write:
//
//	config := LifecycleConfig{}
//	config.SetCreateBeforeDestroy(true)
func (c *LifecycleConfig) SetCreateBeforeDestroy(b bool) {
	c.CreateBeforeDestroy = &b
}

// SetPreventDestroy is a simple helper to avoid having to create a boolean variable and then pass the pointer to it
//
// E.g: instead of writing
//
//	preventDestroy = false
//	config := LifecycleConfig{PreventDestroy: &preventDestroy}
//
// Simply write:
//
//	config := LifecycleConfig{}
//	config.SetPreventDestroy(true)
func (c *LifecycleConfig) SetPreventDestroy(b bool) {
	c.PreventDestroy = &b
}

// LifecycleCondition is used for Precondition and Postcondition property of LifecycleConfig
// It's basically a wrapper for terraform lifecycle pre- and post-conditions.
type LifecycleCondition struct {
	Condition    string
	ErrorMessage string
}

// Lifecycle adds an empty line and the 'lifecycle' terraform directive and then append provided lifecycle attributes.
func (sig *BlockSignature) Lifecycle(config LifecycleConfig) {
	lifecycleSig := NewEmptySignature("lifecycle")

	appendLifecycleBoolAttribute(lifecycleSig, "create_before_destroy", config.CreateBeforeDestroy)
	appendLifecycleBoolAttribute(lifecycleSig, "prevent_destroy", config.PreventDestroy)

	if config.IgnoreChanges != nil {
		lifecycleSig.AppendAttribute("ignore_changes", *tokens.NewIdentListValue(config.IgnoreChanges))
	}

	appendLifecycleConditionBlock(lifecycleSig, "precondition", config.Precondition)
	appendLifecycleConditionBlock(lifecycleSig, "postcondition", config.Postcondition)

	sig.AppendEmptyLine()
	sig.AppendChild(lifecycleSig)
}

/** Private **/

func appendLifecycleConditionBlock(lifecycleSig *BlockSignature, name string, lcCond *LifecycleCondition) {
	if lcCond == nil {
		return
	}

	cond := NewEmptySignature(name)

	cond.AppendAttribute("condition", *tokens.NewIdentValue(lcCond.Condition))
	cond.AppendAttribute("error_message", cty.StringVal(lcCond.ErrorMessage))

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
