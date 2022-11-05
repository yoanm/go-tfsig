package tokens

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func Generate(valuePtr *cty.Value) hclwrite.Tokens {
	var value cty.Value
	if valuePtr != nil {
		value = *valuePtr
		valType := value.Type()
		switch {
		case valType == cty.NilType:
			// Do nothing, let hclwrite.TokensForValue do the job
		case valType.IsListType() || valType.IsSetType() || valType.IsTupleType():
			if ContainsCapsule(valuePtr) {
				// Generate new element token
				newElements := make([]hclwrite.Tokens, value.LengthInt())
				currentIndex := 0
				for it := value.ElementIterator(); it.Next(); {
					_, eVal := it.Element()
					newElements[currentIndex] = Generate(&eVal)
					currentIndex++
				}

				return GenerateFromIterable(valType, newElements)
			}

		case valType.IsMapType() || valType.IsObjectType():
			if ContainsCapsule(valuePtr) {
				// Generate new element token
				newElements := make([]hclwrite.Tokens, value.LengthInt())
				currentIndex := 0
				for it := value.ElementIterator(); it.Next(); {
					eKey, eVal := it.Element()
					newElements[currentIndex] = Generate(&eVal).BuildTokens(NewEqualTokens()).BuildTokens(Generate(&eKey))
					currentIndex++
				}

				return GenerateFromIterable(valType, newElements)
			}
		case IsCapsule(value.Type()):
			return FromValue(value)
		}
	} else {
		return hclwrite.Tokens{}
	}

	return hclwrite.TokensForValue(value)
}

// GenerateFromIterable will panic if provided type is not an iterable type
func GenerateFromIterable(t cty.Type, elements []hclwrite.Tokens) hclwrite.Tokens {
	var emptyCollectionValue cty.Value
	switch {
	case t.IsListType():
		emptyCollectionValue = cty.ListValEmpty(t.ElementType())
	case t.IsSetType():
		emptyCollectionValue = cty.SetValEmpty(t.ElementType())
	case t.IsTupleType():
		emptyCollectionValue = cty.EmptyTupleVal
	case t.IsMapType():
		emptyCollectionValue = cty.MapValEmpty(t.ElementType())
	case t.IsObjectType():
		emptyCollectionValue = cty.EmptyObjectVal
	default:
		panic(fmt.Sprintf("expected a collection type but got %s", t.GoString()))
	}

	return MergeIterableAndGenerate(emptyCollectionValue, elements)
}

// MergeIterableAndGenerate will panic if provided value is not iterable
func MergeIterableAndGenerate(collection cty.Value, newElements []hclwrite.Tokens) hclwrite.Tokens {
	tokensStart, existingElements, tokensEnd := SplitIterable(collection)

	newTokens := existingElements.BuildTokens(tokensStart)

	if len(newElements) > 0 {
		collectionType := collection.Type()
		var separator hclwrite.Tokens
		addSeparator := true
		if collectionType.IsListType() || collectionType.IsSetType() || collectionType.IsTupleType() {
			// Separate elements with a comma for list/set and tuple
			separator = NewCommaTokens()
			addSeparator = len(existingElements) > 0
		} else {
			// Separate elements with a new line for Objects and Maps
			separator = NewLineTokens()
			// Objects and Maps already have a trailing new line if not empty
			// => separator must be added only after a new element is added
		}

		for _, elem := range newElements {
			if addSeparator {
				newTokens = separator.BuildTokens(newTokens)
			}
			newTokens = elem.BuildTokens(newTokens)
			addSeparator = true
		}
		if addSeparator && (collectionType.IsMapType() || collectionType.IsObjectType()) {
			// Object and map have a trailing separator
			newTokens = separator.BuildTokens(newTokens)
		}
	}

	return tokensEnd.BuildTokens(newTokens)
}

// SplitIterable will panic if provided value is not iterable
func SplitIterable(collection cty.Value) (tokensStart hclwrite.Tokens, elements hclwrite.Tokens, tokensEnd hclwrite.Tokens) {
	if !collection.CanIterateElements() {
		panic(fmt.Sprintf("expected an iterable type but got %s", collection.Type().GoString()))
	}

	tokens := hclwrite.TokensForValue(collection)

	var start, elems, end hclwrite.Tokens
	startFound, endFound := false, false

	for _, token := range tokens {
		if token.Type == hclsyntax.TokenCBrack || token.Type == hclsyntax.TokenCBrace {
			endFound = true
		}

		if endFound {
			end = append(end, token)
		} else if startFound {
			elems = append(elems, token)
		} else {
			start = append(start, token)
		}

		if token.Type == hclsyntax.TokenOBrack || token.Type == hclsyntax.TokenOBrace {
			startFound = true
		}
	}

	return start, elems, end
}
