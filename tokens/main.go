package tokens

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const (
	HclwriteTokensCtyTypeName = "cty.CapsuleVal(hclwrite.Tokens)"
)

var hclwriteTokensCtyType cty.Type

func init() {
	hclwriteTokensCtyType = cty.Capsule(HclwriteTokensCtyTypeName, reflect.TypeOf(hclwrite.Tokens{}))
}

func NewIdentValue(s string) *cty.Value {
	val := ToValue(NewIdentTokens(s))

	return &val
}

func NewIdentListValue(list []string) *cty.Value {
	if list == nil {
		return nil
	}

	listLength := len(list)
	val := cty.ListValEmpty(hclwriteTokensCtyType)
	if listLength > 0 {
		newList := make([]cty.Value, listLength)
		for i, s := range list {
			newList[i] = ToValue(NewIdentTokens(s))
		}

		val = cty.ListVal(newList)
	}

	return &val
}

func ToValue(tokens hclwrite.Tokens) cty.Value {
	return cty.CapsuleVal(hclwriteTokensCtyType, &tokens)
}

func FromValue(v cty.Value) (newTokens hclwrite.Tokens) {
	if err := gocty.FromCtyValue(v, &newTokens); err != nil {
		panic(fmt.Sprintf("error during conversion from cty.Value to hclwrite.Tokens: %s", err))
	}

	return newTokens
}
