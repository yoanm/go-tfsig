/*
Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident)

It also provides an easy way to encapsulate hclwrite tokens into a cty.Value and a function (`Generate()`)
to manage those type of value
*/
package tokens

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// HclwriteTokensCtyTypeName is the friendly cty name for the capsule encapsulating `hclwrite.Tokens`.
//
//nolint:gosec // disabled as linter thinks it's an hardcoded credential
const HclwriteTokensCtyTypeName = "cty.CapsuleVal(hclwrite.Tokens)"

//nolint:gochecknoglobals // Better to keep it as **internal** global var than define it each time
var hclwriteTokensCtyType = cty.Capsule(HclwriteTokensCtyTypeName, reflect.TypeOf(hclwrite.Tokens{}))

// NewIdentValue takes a string which should be considered as 'ident' token and converts it
// to a special `cty.Value` capsule.
func NewIdentValue(s string) *cty.Value {
	val := ToValue(NewIdentTokens(s))

	return &val
}

// NewIdentListValue takes a list of string which should be all considered as 'ident' tokens
// and converts them into a cty list containing special `cty.Value` capsule.
func NewIdentListValue(list []string) *cty.Value {
	if list == nil {
		return nil
	}

	val := cty.ListValEmpty(hclwriteTokensCtyType)

	listLength := len(list)
	if listLength > 0 {
		newList := make([]cty.Value, listLength)
		for i, s := range list {
			newList[i] = ToValue(NewIdentTokens(s))
		}

		val = cty.ListVal(newList)
	}

	return &val
}

// ToValue takes `hclwrite.Tokens` value and converts it to special `cty.Value` capsule.
func ToValue(tokens hclwrite.Tokens) cty.Value {
	return cty.CapsuleVal(hclwriteTokensCtyType, &tokens)
}

// FromValue takes a `cty.Value` and extract the `hclwrite.Tokens` from it.
//
// It panics if the provided value is not a special `cty.Value` capsule.
func FromValue(v cty.Value) hclwrite.Tokens {
	newTokens := hclwrite.Tokens{}
	if err := gocty.FromCtyValue(v, &newTokens); err != nil {
		panic(fmt.Sprintf("error during conversion from cty.Value to hclwrite.Tokens: %s", err))
	}

	return newTokens
}
