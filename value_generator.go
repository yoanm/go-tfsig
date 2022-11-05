package tfsig

import (
	"fmt"
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/tfsig/tokens"
)

type ValueGenerator struct {
	matcher IdentTokenMatcher
}

func NewValueGenerator(identPrefixList ...string) ValueGenerator {
	return NewValueGeneratorWith(NewIdentTokenMatcher(identPrefixList...))
}

func NewValueGeneratorWith(matcher IdentTokenMatcher) ValueGenerator {
	return ValueGenerator{matcher: matcher}
}

func (g *ValueGenerator) ToIdent(s *string) *cty.Value {
	if s == nil {
		return nil
	}

	return tokens.NewIdentValue(*s)
}

func (g *ValueGenerator) ToIdentList(list *[]string) *cty.Value {
	if list == nil {
		return nil
	}

	return tokens.NewIdentListValue(*list)
}

func (g *ValueGenerator) ToString(s *string) *cty.Value {
	return g.FromString(s, cty.String)
}
func (g *ValueGenerator) ToBool(s *string) *cty.Value {
	return g.FromString(s, cty.Bool)
}
func (g *ValueGenerator) ToNumber(s *string) *cty.Value {
	return g.FromString(s, cty.Number)
}
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
