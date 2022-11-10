# tokens

Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident)

It also provides an easy way to encapsulate hclwrite tokens into a cty.Value and a function (`Generate()`) to manage those type of value

## Constants

```golang
const (
    // HclwriteTokensCtyTypeName is the friendly cty name for the capsule encapsulating `hclwrite.Tokens`
    HclwriteTokensCtyTypeName = "cty.CapsuleVal(hclwrite.Tokens)"
)
```

## Functions

### func [ContainsCapsule](/tokens/token_capsule.go#L14)

`func ContainsCapsule(valPtr *cty.Value) bool`

ContainsCapsule will deep check if provided value contains a special capsule encapsulating `hclwrite.Tokens`
(and therefore requires special process to de-encapsulate it)

### func [FromValue](/tokens/main.go#L64)

`func FromValue(v cty.Value) (newTokens hclwrite.Tokens)`

FromValue takes a `cty.Value` and extract the `hclwrite.Tokens` from it.

It panics if the provided valud is not a special `cty.Value` capsule

### func [Generate](/tokens/generator.go#L14)

`func Generate(valuePtr *cty.Value) hclwrite.Tokens`

Generate converts a `cty.Value` to `hclwrite.Tokens`

It takes care of special `cty.Value` capsule encapsulating `hclwrite.Tokens`

```golang
listOfCapsule, err := gocty.ToCtyValue(
    []cty.Value{
        *NewIdentValue("value1"),
        *NewIdentValue("value2"),
    },
    cty.List(hclwriteTokensCtyType),
)
if err != nil {
    panic(err)
}

setOfCapsule, err := gocty.ToCtyValue(
    []cty.Value{
        *NewIdentValue("value1"),
        *NewIdentValue("value2"),
    },
    cty.Set(hclwriteTokensCtyType),
)
if err != nil {
    panic(err)
}

objectWithCapsule, err := gocty.ToCtyValue(
    map[string]cty.Value{
        "A": *NewIdentValue("A_value"),
        "B": cty.StringVal("B_value"),
    },
    cty.Object(
        map[string]cty.Type{
            "A": hclwriteTokensCtyType,
            "B": cty.String,
        },
    ),
)
if err != nil {
    panic(err)
}

mapOfCapsule, err := gocty.ToCtyValue(
    map[string]cty.Value{
        "A": *NewIdentValue("A_value"),
        "B": *NewIdentValue("B_value"),
    },
    cty.Map(hclwriteTokensCtyType),
)
if err != nil {
    panic(err)
}

tupleWithCapsule, err := gocty.ToCtyValue(
    []cty.Value{
        cty.StringVal("A_value"),
        *NewIdentValue("B_value"),
        cty.NumberIntVal(2),
    },
    cty.Tuple([]cty.Type{cty.String, hclwriteTokensCtyType, cty.Number}),
)
if err != nil {
    panic(err)
}
stringVal := cty.StringVal("TeSt")
numberIntVal := cty.NumberIntVal(12)
numberFloatVal := cty.NumberFloatVal(-12.23)
boolVal := cty.BoolVal(false)

fmt.Printf("Null: %#v\n", string(Generate(&cty.NilVal).Bytes()))
fmt.Printf("Ident: %#v\n", string(Generate(NewIdentValue("TeSt")).Bytes()))
fmt.Printf("String: %#v\n", string(Generate(&stringVal).Bytes()))
fmt.Printf("Positive number: %#v\n", string(Generate(&numberIntVal).Bytes()))
fmt.Printf("Negative number: %#v\n", string(Generate(&numberFloatVal).Bytes()))
fmt.Printf("Boolean: %#v\n", string(Generate(&boolVal).Bytes()))
fmt.Printf("List of capsule: %#v\n", string(Generate(&listOfCapsule).Bytes()))
fmt.Printf("Set of capsule: %#v\n", string(Generate(&setOfCapsule).Bytes()))
fmt.Printf("Object with capsule: %#v\n", string(Generate(&objectWithCapsule).Bytes()))
fmt.Printf("Map of capsule: %#v\n", string(Generate(&mapOfCapsule).Bytes()))
fmt.Printf("Tuple with capsule: %#v\n", string(Generate(&tupleWithCapsule).Bytes()))
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

### func [GenerateFromIterable](/tokens/generator.go#L62)

`func GenerateFromIterable(elements []hclwrite.Tokens, t cty.Type) hclwrite.Tokens`

GenerateFromIterable takes a list of `hclwrite.Tokens` and create related `hclwrite.Tokens` based on the provided `cty.Type`

It panics if provided type is not an iterable type

### func [IsCapsuleType](/tokens/token_capsule.go#L8)

`func IsCapsuleType(t cty.Type) bool`

IsCapsuleType returns true if provided `cty.Type` is a special capsule encapsulating `hclwrite.Tokens`

### func [MergeIterableAndGenerate](/tokens/generator.go#L85)

`func MergeIterableAndGenerate(collection cty.Value, newElements []hclwrite.Tokens) hclwrite.Tokens`

MergeIterableAndGenerate takes a `cty.Value` collection, append new elements and convert the result to related `hclwrite.Tokens`

It panics if provided collection is not iterable

### func [NewCommaToken](/tokens/token.go#L14)

`func NewCommaToken() *hclwrite.Token`

NewCommaToken returns a `hclwrite.Token` with `hclsyntax.TokenComma` type

### func [NewCommaTokens](/tokens/tokens.go#L17)

`func NewCommaTokens() hclwrite.Tokens`

NewCommaTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenComma` type

See also `NewCommaToken()`

### func [NewEqualToken](/tokens/token.go#L19)

`func NewEqualToken() *hclwrite.Token`

NewEqualToken returns a `hclwrite.Token` with `hclsyntax.TokenEqual` type

### func [NewEqualTokens](/tokens/tokens.go#L24)

`func NewEqualTokens() hclwrite.Tokens`

NewEqualTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenEqual` type

See also `NewEqualToken()`

### func [NewIdentListValue](/tokens/main.go#L37)

`func NewIdentListValue(list []string) *cty.Value`

NewIdentListValue tales a list of string which should be all considered as 'ident' tokens
and converts them into a cty list containing special `cty.Value` capsule

```golang
identListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}
value := NewIdentListValue(identListStringValue)

// ... Later in the code
hclFile := hclwrite.NewEmptyFile()
hclFile.Body().SetAttributeRaw("attr", Generate(value))

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
attr = [explicit_ident_item.foo, explicit_ident_item.bar]
```

### func [NewIdentToken](/tokens/token.go#L9)

`func NewIdentToken(b []byte) *hclwrite.Token`

NewIdentToken returns a `hclwrite.Token` with `hclsyntax.TokenIdent` type encapsulating provided bytes

### func [NewIdentTokens](/tokens/tokens.go#L10)

`func NewIdentTokens(s string) hclwrite.Tokens`

NewIdentTokens takes a string and convert it to `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenIdent` type

See also `NewIdentToken()`

### func [NewIdentValue](/tokens/main.go#L29)

`func NewIdentValue(s string) *cty.Value`

NewIdentValue takes a string which should be considered as 'ident' token and converts it to a special `cty.Value` capsule

```golang
identStringValue := "explicit_ident.foo"
value := NewIdentValue(identStringValue)

// ... Later in the code
hclFile := hclwrite.NewEmptyFile()
hclFile.Body().SetAttributeRaw("attr", Generate(value))

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
attr = explicit_ident.foo
```

### func [NewLineToken](/tokens/token.go#L24)

`func NewLineToken() *hclwrite.Token`

NewLineToken returns a `hclwrite.Token` with `hclsyntax.TokenNewline` type

### func [NewLineTokens](/tokens/tokens.go#L31)

`func NewLineTokens() hclwrite.Tokens`

NewLineTokens creates a `hclwrite.Tokens` containing a `hclwrite.Token` with `hclsyntax.TokenNewline` type

See also `NewLineToken()`

### func [SplitIterable](/tokens/generator.go#L126)

`func SplitIterable(collection cty.Value) (tokensStart hclwrite.Tokens, elements hclwrite.Tokens, tokensEnd hclwrite.Tokens)`

SplitIterable takes a `cty.Value` collection and returns the start/end tokens and the existing elements

It can be used to later append new elements to the collection (see `MergeIterableAndGenerate()`)

It panics if provided collection is not iterable

```golang
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

start, elements, end := SplitIterable(list)
fmt.Printf(
    "List:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
    string(start.Bytes()),
    string(elements.Bytes()),
    string(end.Bytes()),
)
start, elements, end = SplitIterable(setOfCapsule)
fmt.Printf(
    "Set:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
    string(start.Bytes()),
    string(elements.Bytes()),
    string(end.Bytes()),
)
start, elements, end = SplitIterable(objectWithCapsule)
fmt.Printf(
    "Object:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
    string(start.Bytes()),
    string(elements.Bytes()),
    string(end.Bytes()),
)
start, elements, end = SplitIterable(mapOfCapsule)
fmt.Printf(
    "Map:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
    string(start.Bytes()),
    string(elements.Bytes()),
    string(end.Bytes()),
)
start, elements, end = SplitIterable(tupleWithCapsule)
fmt.Printf(
    "Tuple:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
    string(start.Bytes()),
    string(elements.Bytes()),
    string(end.Bytes()),
)
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

### func [ToValue](/tokens/main.go#L57)

`func ToValue(tokens hclwrite.Tokens) cty.Value`

ToValue takes `hclwrite.Tokens` value and converts it to special `cty.Value` capsule

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
