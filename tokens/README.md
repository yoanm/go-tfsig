# tokens

Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident)

It also provides an easy way to encapsulate hclwrite tokens into a cty.Value and a function (`Generate()`)
to manage those type of value

## Constants

HclwriteTokensCtyTypeName is the friendly cty name for the capsule encapsulating `hclwrite.Tokens`.

```golang
const HclwriteTokensCtyTypeName = "cty.CapsuleVal(hclwrite.Tokens)"
```

## Functions

### func [ContainsCapsule](./token_capsule.go#L14)

`func ContainsCapsule(valPtr *cty.Value) bool`

ContainsCapsule will deep check if provided value contains a special capsule encapsulating `hclwrite.Tokens`
(and therefore requires special process to de-encapsulate it).

### func [FromValue](./main.go#L64)

`func FromValue(v cty.Value) hclwrite.Tokens`

FromValue takes a `cty.Value` and extract the `hclwrite.Tokens` from it.

It panics if the provided value is not a special `cty.Value` capsule.

### func [Generate](./generator.go#L16)

`func Generate(valuePtr *cty.Value) hclwrite.Tokens`

Generate converts a `cty.Value` to `hclwrite.Tokens`

It takes care of special `cty.Value` capsule encapsulating `hclwrite.Tokens`.

```golang
package main

import (
	"fmt"
	"github.com/yoanm/go-tfsig/tokens"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func main() {
	listOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*tokens.NewIdentValue("value1"),
			*tokens.NewIdentValue("value2"),
		},
		cty.List(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*tokens.NewIdentValue("value1"),
			*tokens.NewIdentValue("value2"),
		},
		cty.Set(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	objectWithCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *tokens.NewIdentValue("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Object(
			map[string]cty.Type{
				"A": cty.DynamicPseudoType,
				"B": cty.String,
			},
		),
	)
	if err != nil {
		panic(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *tokens.NewIdentValue("A_value"),
			"B": *tokens.NewIdentValue("B_value"),
		},
		cty.Map(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			*tokens.NewIdentValue("B_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, cty.DynamicPseudoType, cty.Number}),
	)
	if err != nil {
		panic(err)
	}

	stringVal := cty.StringVal("TeSt")
	numberIntVal := cty.NumberIntVal(12)
	numberFloatVal := cty.NumberFloatVal(-12.23)
	boolVal := cty.BoolVal(false)

	fmt.Printf("Null: %#v\n", string(tokens.Generate(&cty.NilVal).Bytes()))
	fmt.Printf("Ident: %#v\n", string(tokens.Generate(tokens.NewIdentValue("TeSt")).Bytes()))
	fmt.Printf("String: %#v\n", string(tokens.Generate(&stringVal).Bytes()))
	fmt.Printf("Positive number: %#v\n", string(tokens.Generate(&numberIntVal).Bytes()))
	fmt.Printf("Negative number: %#v\n", string(tokens.Generate(&numberFloatVal).Bytes()))
	fmt.Printf("Boolean: %#v\n", string(tokens.Generate(&boolVal).Bytes()))
	fmt.Printf("List of capsule: %#v\n", string(tokens.Generate(&listOfCapsule).Bytes()))
	fmt.Printf("Set of capsule: %#v\n", string(tokens.Generate(&setOfCapsule).Bytes()))
	fmt.Printf("Object with capsule: %#v\n", string(tokens.Generate(&objectWithCapsule).Bytes()))
	fmt.Printf("Map of capsule: %#v\n", string(tokens.Generate(&mapOfCapsule).Bytes()))
	fmt.Printf("Tuple with capsule: %#v\n", string(tokens.Generate(&tupleWithCapsule).Bytes()))

}

```

 Output:

```
Null: "null"
Ident: "TeSt"
String: "\"TeSt\""
Positive number: "12"
Negative number: "-12.23"
Boolean: "false"
List of capsule: "[value1,value2]"
Set of capsule: "[value1,value2]"
Object with capsule: "{\n\"A\"=A_value\n\"B\"=\"B_value\"\n}"
Map of capsule: "{\n\"A\"=A_value\n\"B\"=B_value\n}"
Tuple with capsule: "[\"A_value\",B_value,2]"
```

### func [GenerateFromIterable](./generator.go#L44)

`func GenerateFromIterable(elements []hclwrite.Tokens, toType cty.Type) hclwrite.Tokens`

GenerateFromIterable takes a list of `hclwrite.Tokens` and create related `hclwrite.Tokens` based on
the provided `cty.Type`

It panics if provided type is not an iterable type.

### func [IsCapsuleType](./token_capsule.go#L8)

`func IsCapsuleType(t cty.Type) bool`

IsCapsuleType returns true if provided `cty.Type` is a special capsule encapsulating `hclwrite.Tokens`.

### func [MergeIterableAndGenerate](./generator.go#L69)

`func MergeIterableAndGenerate(collection cty.Value, newElements []hclwrite.Tokens) hclwrite.Tokens`

MergeIterableAndGenerate takes a `cty.Value` collection, append new elements and convert the result
to related `hclwrite.Tokens`

It panics if provided collection is not iterable.

### func [NewCommaToken](./token.go#L14)

`func NewCommaToken() *hclwrite.Token`

NewCommaToken returns a `hclwrite.Token` with `hclsyntax.TokenComma` type.

### func [NewCommaTokens](./tokens.go#L18)

`func NewCommaTokens() hclwrite.Tokens`

NewCommaTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenComma` type

See also `NewCommaToken()`.

### func [NewEqualToken](./token.go#L19)

`func NewEqualToken() *hclwrite.Token`

NewEqualToken returns a `hclwrite.Token` with `hclsyntax.TokenEqual` type.

### func [NewEqualTokens](./tokens.go#L25)

`func NewEqualTokens() hclwrite.Tokens`

NewEqualTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenEqual` type

See also `NewEqualToken()`.

### func [NewIdentListValue](./main.go#L36)

`func NewIdentListValue(list []string) *cty.Value`

NewIdentListValue takes a list of string which should be all considered as 'ident' tokens
and converts them into a cty list containing special `cty.Value` capsule.

```golang
package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig/tokens"
)

func main() {
	identListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}
	value := tokens.NewIdentListValue(identListStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", tokens.Generate(value))

	fmt.Println(string(hclFile.Bytes()))
}

```

 Output:

```
attr = [explicit_ident_item.foo, explicit_ident_item.bar]
```

### func [NewIdentToken](./token.go#L9)

`func NewIdentToken(b []byte) *hclwrite.Token`

NewIdentToken returns a `hclwrite.Token` with `hclsyntax.TokenIdent` type encapsulating provided bytes.

### func [NewIdentTokens](./tokens.go#L11)

`func NewIdentTokens(s string) hclwrite.Tokens`

NewIdentTokens takes a string and convert it to `hclwrite.Tokens` containing a `hclwrite.Token`
with `hclsyntax.TokenIdent` type

See also `NewIdentToken()`.

### func [NewIdentValue](./main.go#L28)

`func NewIdentValue(s string) *cty.Value`

NewIdentValue takes a string which should be considered as 'ident' token and converts it
to a special `cty.Value` capsule.

```golang
package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/yoanm/go-tfsig/tokens"
)

func main() {
	identStringValue := "explicit_ident.foo"
	value := tokens.NewIdentValue(identStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", tokens.Generate(value))

	fmt.Println(string(hclFile.Bytes()))
}

```

 Output:

```
attr = explicit_ident.foo
```

### func [NewLineToken](./token.go#L24)

`func NewLineToken() *hclwrite.Token`

NewLineToken returns a `hclwrite.Token` with `hclsyntax.TokenNewline` type.

### func [NewLineTokens](./tokens.go#L32)

`func NewLineTokens() hclwrite.Tokens`

NewLineTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenNewline` type

See also `NewLineToken()`.

### func [SplitIterable](./generator.go#L115)

`func SplitIterable(collection cty.Value) (
    hclwrite.Tokens,
    hclwrite.Tokens,
    hclwrite.Tokens,
)`

SplitIterable takes a `cty.Value` collection and returns the start tokens, the existing elements tokens
and the end tokens

It can be used to later append new elements to the collection (see `MergeIterableAndGenerate()`)

It panics if provided collection is not iterable.

```golang
package main

import (
	"fmt"
	"github.com/yoanm/go-tfsig/tokens"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func main() {
	list, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.List(cty.String),
	)
	if err != nil {
		panic(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.Set(cty.String),
	)
	if err != nil {
		panic(err)
	}

	objectWithCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": cty.NumberIntVal(2),
			"B": cty.StringVal("B_value"),
		},
		cty.Object(
			map[string]cty.Type{
				"A": cty.Number,
				"B": cty.String,
			},
		),
	)
	if err != nil {
		panic(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": cty.StringVal("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Map(cty.String),
	)
	if err != nil {
		panic(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, cty.Number}),
	)
	if err != nil {
		panic(err)
	}

	start, elements, end := tokens.SplitIterable(list)
	fmt.Printf(
		"List:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(setOfCapsule)
	fmt.Printf(
		"Set:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(objectWithCapsule)
	fmt.Printf(
		"Object:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(mapOfCapsule)
	fmt.Printf(
		"Map:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(tupleWithCapsule)
	fmt.Printf(
		"Tuple:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

}

```

 Output:

```
List:
	Start: "["
	Elements: "\"value1\", \"value2\""
	End: "]"
Set:
	Start: "["
	Elements: "\"value1\", \"value2\""
	End: "]"
Object:
	Start: "{"
	Elements: "\n  A = 2\n  B = \"B_value\"\n"
	End: "}"
Map:
	Start: "{"
	Elements: "\n  A = \"A_value\"\n  B = \"B_value\"\n"
	End: "}"
Tuple:
	Start: "["
	Elements: "\"A_value\", 2"
	End: "]"
```

### func [ToValue](./main.go#L57)

`func ToValue(tokens hclwrite.Tokens) cty.Value`

ToValue takes `hclwrite.Tokens` value and converts it to special `cty.Value` capsule.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
