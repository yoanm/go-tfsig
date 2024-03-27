package tokens

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

/** Public **/

// Generate converts a `cty.Value` to `hclwrite.Tokens`
//
// It takes care of special `cty.Value` capsule encapsulating `hclwrite.Tokens`.
func Generate(valuePtr *cty.Value) hclwrite.Tokens {
	var value cty.Value
	if valuePtr != nil {
		value = *valuePtr
		valType := value.Type()

		switch {
		case valType == cty.NilType:
			// Do nothing, let `hclwrite.TokensForValue()` do the job
		case valType.IsListType() || valType.IsSetType() || valType.IsTupleType() ||
			valType.IsMapType() || valType.IsObjectType():
			if ContainsCapsule(valuePtr) {
				return generateCapsuleForCollection(valType, value)
			}
		case IsCapsuleType(value.Type()):
			return FromValue(value)
		}
	} else {
		return hclwrite.Tokens{}
	}

	return hclwrite.TokensForValue(value)
}

// GenerateFromIterable takes a list of `hclwrite.Tokens` and create related `hclwrite.Tokens` based on
// the provided `cty.Type`
//
// It panics if provided type is not an iterable type.
func GenerateFromIterable(elements []hclwrite.Tokens, toType cty.Type) hclwrite.Tokens {
	var emptyCollectionValue cty.Value

	switch {
	case toType.IsListType():
		emptyCollectionValue = cty.ListValEmpty(toType.ElementType())
	case toType.IsSetType():
		emptyCollectionValue = cty.SetValEmpty(toType.ElementType())
	case toType.IsTupleType():
		emptyCollectionValue = cty.EmptyTupleVal
	case toType.IsMapType():
		emptyCollectionValue = cty.MapValEmpty(toType.ElementType())
	case toType.IsObjectType():
		emptyCollectionValue = cty.EmptyObjectVal
	default:
		panic("expected a collection type but got " + toType.GoString())
	}

	return MergeIterableAndGenerate(emptyCollectionValue, elements)
}

// MergeIterableAndGenerate takes a `cty.Value` collection, append new elements and convert the result
// to related `hclwrite.Tokens`
//
// It panics if provided collection is not iterable.
func MergeIterableAndGenerate(collection cty.Value, newElements []hclwrite.Tokens) hclwrite.Tokens {
	tokensStart, existingElements, tokensEnd := SplitIterable(collection)

	newTokens := existingElements.BuildTokens(tokensStart)

	if len(newElements) > 0 {
		var separator hclwrite.Tokens

		addSeparator := true

		collectionType := collection.Type()
		if collectionType.IsListType() || collectionType.IsSetType() || collectionType.IsTupleType() {
			// Separate elements with a comma for list/set and tuple
			separator = NewCommaTokens()
			addSeparator = len(existingElements) > 0
		} else {
			// Separate elements with a new line for Objects and Maps
			// Objects and Maps already have a trailing new line if not empty
			// => separator must be added only after a new element is added
			separator = NewLineTokens()
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

// SplitIterable takes a `cty.Value` collection and returns the start tokens, the existing elements tokens
// and the end tokens
//
// It can be used to later append new elements to the collection (see `MergeIterableAndGenerate()`)
//
// It panics if provided collection is not iterable.
func SplitIterable(collection cty.Value) (
	/* tokensStart */ hclwrite.Tokens,
	/* elements */ hclwrite.Tokens,
	/* tokensEnd */ hclwrite.Tokens,
) {
	if !collection.CanIterateElements() {
		panic("expected an iterable type but got " + collection.Type().GoString())
	}

	var start, elems, end hclwrite.Tokens

	startFound, endFound := false, false

	tokens := hclwrite.TokensForValue(collection)
	for _, token := range tokens {
		if token.Type == hclsyntax.TokenCBrack || token.Type == hclsyntax.TokenCBrace {
			endFound = true
		}

		switch {
		case endFound:
			end = append(end, token)
		case startFound:
			elems = append(elems, token)
		default:
			start = append(start, token)
		}

		if token.Type == hclsyntax.TokenOBrack || token.Type == hclsyntax.TokenOBrace {
			startFound = true
		}
	}

	return start, elems, end
}

/** Private **/

func generateCapsuleForCollection(valType cty.Type, value cty.Value) hclwrite.Tokens {
	isMapOrObjectType := valType.IsMapType() || valType.IsObjectType()
	// Generate new element token
	newElements := make([]hclwrite.Tokens, value.LengthInt())
	currentIndex := 0

	for it := value.ElementIterator(); it.Next(); {
		var tokens hclwrite.Tokens

		eKey, eVal := it.Element()
		if isMapOrObjectType {
			tokens = Generate(&eVal).BuildTokens(NewEqualTokens()).BuildTokens(Generate(&eKey))
		} else {
			tokens = Generate(&eVal)
		}

		newElements[currentIndex] = tokens
		currentIndex++
	}

	return GenerateFromIterable(newElements, valType)
}
