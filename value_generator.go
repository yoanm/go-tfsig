package tfsig

import (
	"fmt"
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/tokens"
)

// ValueGenerator is able to detect "ident" tokens and convert them into a special cty capsule.
// Capsule will then be converted to `hclwrite.tokens`
// It allows to write values like `var.my_var`, `locals.my_local` or `data.res_name.val_name` without any quotes
type ValueGenerator struct {
	matcher IdentTokenMatcherInterface
}

// NewValueGenerator returns a new ValueGenerator with the default 'ident' tokens matcher augmented with provided list
// of token to consider as 'ident' tokens
func NewValueGenerator(identPrefixList ...string) ValueGenerator {
	return NewValueGeneratorWith(NewIdentTokenMatcher(identPrefixList...))
}

// NewValueGeneratorWith returns a new ValueGenerator with the provided matcher
func NewValueGeneratorWith(matcher IdentTokenMatcherInterface) ValueGenerator {
	return ValueGenerator{matcher: matcher}
}

// ToIdent converts a string to a special `cty.Value` capsule holding `hclwrite.tokens`
func (g *ValueGenerator) ToIdent(s *string) *cty.Value {
	if s == nil {
		return nil
	}

	return tokens.NewIdentValue(*s)
}

// ToIdentList converts a list of string to `cty.Value` list containing capsules holding `hclwrite.tokens`
func (g *ValueGenerator) ToIdentList(list *[]string) *cty.Value {
	if list == nil {
		return nil
	}

	return tokens.NewIdentListValue(*list)
}

// ToString convert a string to `cty.Value` string which will be rendered as quoted string by terraform HCL
// If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`
func (g *ValueGenerator) ToString(s *string) *cty.Value {
	return g.FromString(s, cty.String)
}

// ToBool convert a string to `cty.Value` boolean which will be rendered as true or false value by terraform HCL
// If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`
func (g *ValueGenerator) ToBool(s *string) *cty.Value {
	return g.FromString(s, cty.Bool)
}

// ToNumber convert a string to `cty.Value` number which will be rendered as numeric value by terraform HCL
// If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`
func (g *ValueGenerator) ToNumber(s *string) *cty.Value {
	return g.FromString(s, cty.Number)
}

// ToStringList convert a string list to `cty.Value` string list which will be rendered as quoted string list by terraform HCL
// If a provided string item is actually an 'ident' token, `cty.Value` item will be a capsule holding `hclwrite.tokens`
func (g *ValueGenerator) ToStringList(list *[]string) *cty.Value {
	if list == nil {
		return nil
	}
	listLength := len(*list)
	val := cty.EmptyTupleVal
	if listLength > 0 {
		newList := make([]cty.Value, listLength)
		for i, rawValue := range *list {
			newList[i] = *g.FromString(&rawValue, cty.String)
		}

		val = cty.TupleVal(newList)
	}

	return &val
}

// FromString convert a string to `cty.Value` of the provided type
// If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`
func (g *ValueGenerator) FromString(s *string, t cty.Type) *cty.Value {
	if s == nil {
		return nil
	}
	if g.matcher.IsIdentToken(*s) {
		return g.ToIdent(s)
	}

	switch t {
	case cty.String:
		val := cty.StringVal(*s)

		return &val
	case cty.Bool:
		val := cty.BoolVal(*s == "1" || strings.ToLower(*s) == "true")

		return &val
	case cty.Number:
		val := cty.MustParseNumberVal(*s)

		return &val
	default:
		panic(fmt.Sprintf("Unable to convert \"%s\" to a %s", *s, t.FriendlyName()))
	}
}
