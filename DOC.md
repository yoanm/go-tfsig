# tfsig

Package tfsig is a wrapper for Terraform HCL language (`hclwrite`).

It provides ability to generate block signature which are way easier to manipulate and alter than hclwrite.tokens type

## Functions

### func [AppendBlockIfNotNil](/utils.go#L26)

`func AppendBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block)`

AppendBlockIfNotNil appends the provided block to the provided body only if block is not nil

It simply avoids an `if` in your code.

```golang
hclFile := hclwrite.NewEmptyFile()
evenFn := func(i int) *hclwrite.Block {
    if i%2 == 0 {
        return nil
    }

    return NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
}

AppendBlockIfNotNil(hclFile.Body(), evenFn(0))
AppendBlockIfNotNil(hclFile.Body(), evenFn(1))
AppendBlockIfNotNil(hclFile.Body(), evenFn(2))
AppendBlockIfNotNil(hclFile.Body(), evenFn(3))
AppendBlockIfNotNil(hclFile.Body(), evenFn(4))

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
block1 {
}
block3 {
}
```

### func [AppendNewLineAndBlockIfNotNil](/utils.go#L35)

`func AppendNewLineAndBlockIfNotNil(body *hclwrite.Body, block *hclwrite.Block)`

AppendNewLineAndBlockIfNotNil appends an empty line followed by provided block to the provided body only if block is not nil

It simply avoids an `if` in your code.

```golang
hclFile := hclwrite.NewEmptyFile()
oddFn := func(i int) *hclwrite.Block {
    if i%2 != 0 {
        return nil
    }

    return NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
}

AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(0))
AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(1))
AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(2))
AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(3))
AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(4))

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
block0 {
}

block2 {
}

block4 {
}
```

### func [ToTerraformIdentifier](/utils.go#L45)

`func ToTerraformIdentifier(s string) string`

ToTerraformIdentifier converts a string to a terraform identifier, by converting not allowed characters to `-`

And if provided value starts with a character not allowed as first character, it replaces it by `_`.

```golang
fmt.Printf("a_valid-identifier becomes %s\n", ToTerraformIdentifier("a_valid-identifier"))
fmt.Printf(".github becomes %s\n", ToTerraformIdentifier(".github"))
fmt.Printf("an.identifier becomes %s\n", ToTerraformIdentifier("an.identifier"))
fmt.Printf("0id becomes %s\n", ToTerraformIdentifier("0id"))
```

 Output:

```
a_valid-identifier becomes a_valid-identifier
.github becomes _github
an.identifier becomes an-identifier
0id becomes _id
```

## Types

### type [BlockSignature](/block_signature.go#L38)

`type BlockSignature struct { ... }`

BlockSignature is basically a wrapper to HCL blocks
It holds a type, the block labels and its elements.

#### func [NewEmptyResource](/block_signature.go#L32)

`func NewEmptyResource(name, id string, labels ...string) *BlockSignature`

NewEmptyResource returns a BlockSignature pointer with "resource" type and filled with provided labels.

#### func [NewEmptySignature](/block_signature.go#L27)

`func NewEmptySignature(name string, labels ...string) *BlockSignature`

NewEmptySignature returns a BlockSignature pointer filled with provided labels.

#### func [NewSignature](/block_signature.go#L18)

`func NewSignature(name string, labels []string, elements BodyElements) *BlockSignature`

NewSignature returns a BlockSignature pointer filled with provided labels and elements.

#### func (*BlockSignature) [AppendAttrIfNotNil](/block_signature_extra.go#L14)

`func (s *BlockSignature) AppendAttrIfNotNil(attrName string, v *cty.Value)`

AppendAttrIfNotNil appends the provided attribute only if not nil

It simply avoids an `if` in your code.

#### func (*BlockSignature) [AppendAttribute](/block_signature.go#L70)

`func (s *BlockSignature) AppendAttribute(name string, value cty.Value)`

AppendAttribute appends an attribute to the block.

#### func (*BlockSignature) [AppendChild](/block_signature.go#L75)

`func (s *BlockSignature) AppendChild(child *BlockSignature)`

AppendChild appends a child block to the block.

#### func (*BlockSignature) [AppendChildIfNotNil](/block_signature_extra.go#L24)

`func (s *BlockSignature) AppendChildIfNotNil(child *BlockSignature)`

AppendChildIfNotNil appends the provided child only if not nil. And in case there is existing elements,
it prepends an empty line

It simply avoids an `if` in your code.

#### func (*BlockSignature) [AppendElement](/block_signature.go#L65)

`func (s *BlockSignature) AppendElement(element BodyElement)`

AppendElement appends an element to the block.

#### func (*BlockSignature) [AppendEmptyLine](/block_signature.go#L80)

`func (s *BlockSignature) AppendEmptyLine()`

AppendEmptyLine appends an empty line to the block.

#### func (*BlockSignature) [Build](/block_signature.go#L85)

`func (s *BlockSignature) Build() *hclwrite.Block`

Build creates a `hclwrite.Block` and appends block's elements to it.

#### func (*BlockSignature) [BuildTokens](/block_signature.go#L94)

`func (s *BlockSignature) BuildTokens() (tks hclwrite.Tokens)`

BuildTokens builds the block signature as `hclwrite.Tokens`.

#### func (*BlockSignature) [DependsOn](/block_signature_extra.go#L35)

`func (s *BlockSignature) DependsOn(idList []string)`

DependsOn adds an empty line and the 'depends_on' terraform directive with provided id list.

```golang
// resource with 'depends_on' directive
sig := NewEmptyResource("res_name", "res_id")
sig.AppendAttribute("attribute1", cty.StringVal("value1"))
sig.DependsOn([]string{"another_res.res_id", "another_another_res.res_id"})

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "res_name" "res_id" {
  attribute1 = "value1"

  depends_on = [another_res.res_id, another_another_res.res_id]
}
```

#### func (*BlockSignature) [GetElements](/block_signature.go#L55)

`func (s *BlockSignature) GetElements() BodyElements`

GetElements returns all elements attached to the block.

#### func (*BlockSignature) [GetLabels](/block_signature.go#L50)

`func (s *BlockSignature) GetLabels() []string`

GetLabels returns labels attached to the block.

#### func (*BlockSignature) [GetType](/block_signature.go#L45)

`func (s *BlockSignature) GetType() string`

GetType returns the type of the block.

#### func (*BlockSignature) [Lifecycle](/block_signature_extra.go#L89)

`func (s *BlockSignature) Lifecycle(config LifecycleConfig)`

Lifecycle adds an empty line and the 'lifecycle' terraform directive and then append provided lifecycle attributes.

```golang
// resource with 'lifecycle' directive
sig := NewEmptyResource("res_name", "res_id")
sig.AppendAttribute("attribute1", cty.StringVal("value1"))
config := LifecycleConfig{}
config.SetCreateBeforeDestroy(true)
config.SetPreventDestroy(false)
sig.Lifecycle(config)

sig2 := NewEmptyResource("res2_name", "res2_id")
sig2.AppendAttribute("attribute1", cty.StringVal("value1"))
config2 := LifecycleConfig{
    IgnoreChanges: []string{"attribute1"},
    Postcondition: &LifecycleCondition{
        condition:    "res_name.res_id.attribute1 != \"value1\"",
        errorMessage: "res_name.res_id.attribute1 must equal \"value1\"",
    },
}
sig2.Lifecycle(config2)

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())
hclFile.Body().AppendBlock(sig2.Build())

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "res_name" "res_id" {
  attribute1 = "value1"

  lifecycle {
    create_before_destroy = true
    prevent_destroy       = false
  }
}
resource "res2_name" "res2_id" {
  attribute1 = "value1"

  lifecycle {
    ignore_changes = [attribute1]
    postcondition {
      condition     = res_name.res_id.attribute1 != "value1"
      error_message = "res_name.res_id.attribute1 must equal \"value1\""
    }
  }
}
```

#### func (*BlockSignature) [SetElements](/block_signature.go#L60)

`func (s *BlockSignature) SetElements(elements BodyElements)`

SetElements overrides existing elements by provided ones.

### type [BodyElement](/body_element.go#L24)

`type BodyElement struct { ... }`

BodyElement is a wrapper for more or less anything that can be appended to a BlockSignature.

#### func [NewBodyAttribute](/body_element.go#L14)

`func NewBodyAttribute(name string, attr cty.Value) BodyElement`

NewBodyAttribute returns an Attribute BodyElement.

#### func [NewBodyBlock](/body_element.go#L9)

`func NewBodyBlock(block *BlockSignature) BodyElement`

NewBodyBlock returns a Block BodyElement.

#### func [NewBodyEmptyLine](/body_element.go#L19)

`func NewBodyEmptyLine() BodyElement`

NewBodyEmptyLine returns an empty line BodyElement.

#### func (BodyElement) [Build](/body_element.go#L77)

`func (e BodyElement) Build() *hclwrite.Block`

Build convert the current BodyElement into a `hclwrite.Block`

it panics if BodyElement is not a block (use `IsBodyBlock()` first).

#### func (BodyElement) [GetBodyAttribute](/body_element.go#L57)

`func (e BodyElement) GetBodyAttribute() *cty.Value`

GetBodyAttribute returns the value of the attribute behind the BodyElement

It panics if BodyElement is not an attribute (use `IsBodyAttribute()` first).

#### func (BodyElement) [GetBodyBlock](/body_element.go#L67)

`func (e BodyElement) GetBodyBlock() *BlockSignature`

GetBodyBlock returns the block behind the BodyElement

it panics if BodyElement is not a block (use `IsBodyBlock()` first).

#### func (BodyElement) [GetName](/body_element.go#L35)

`func (e BodyElement) GetName() string`

GetName returns the name of the BodyElement.

#### func (BodyElement) [IsBodyAttribute](/body_element.go#L45)

`func (e BodyElement) IsBodyAttribute() bool`

IsBodyAttribute returns true if the BodyElement is an attribute.

#### func (BodyElement) [IsBodyBlock](/body_element.go#L40)

`func (e BodyElement) IsBodyBlock() bool`

IsBodyBlock returns true if the BodyElement is a block.

#### func (BodyElement) [IsBodyEmptyLine](/body_element.go#L50)

`func (e BodyElement) IsBodyEmptyLine() bool`

IsBodyEmptyLine returns true if the BodyElement is an empty line.

### type [BodyElements](/body_element.go#L32)

`type BodyElements []BodyElement`

BodyElements is a simple wrapper for a list of BodyElement.

### type [IdentTokenMatcher](/ident_token_matcher.go#L32)

`type IdentTokenMatcher struct { ... }`

IdentTokenMatcher is a simple implementation for IdentTokenMatcherInterface.

#### func [NewIdentTokenMatcher](/ident_token_matcher.go#L22)

`func NewIdentTokenMatcher(prefixList ...string) IdentTokenMatcher`

NewIdentTokenMatcher returns an instance of IdentTokenMatcher with provided list of prefix to consider as 'ident' tokens

`local.`, `var.` and `data.` tokens will be considered as 'ident' tokens by default.

#### func (IdentTokenMatcher) [IsIdentToken](/ident_token_matcher.go#L38)

`func (m IdentTokenMatcher) IsIdentToken(s string) bool`

IsIdentToken is the implementation for IdentTokenMatcherInterface.

### type [IdentTokenMatcherInterface](/ident_token_matcher.go#L27)

`type IdentTokenMatcherInterface interface { ... }`

IdentTokenMatcherInterface is a simple interface declaring required method to detect an 'ident' token.

### type [LifecycleCondition](/block_signature_extra.go#L83)

`type LifecycleCondition struct { ... }`

LifecycleCondition is used for Precondition and Postcondition property of LifecycleConfig
It's basically a wrapper for terraform lifecycle pre- and post-conditions.

### type [LifecycleConfig](/block_signature_extra.go#L42)

`type LifecycleConfig struct { ... }`

LifecycleConfig is used as argument for `Lifecycle()` method
It's basically a wrapper for terraform `lifecycle` directive.

#### func (*LifecycleConfig) [SetCreateBeforeDestroy](/block_signature_extra.go#L62)

`func (c *LifecycleConfig) SetCreateBeforeDestroy(b bool)`

SetCreateBeforeDestroy is a simple helper to avoid having to create a boolean variable and then pass the pointer to it

E.g: instead of writing

```go
createBeforeDestroy = true
config := LifecycleConfig{CreateBeforeDestroy: &createBeforeDestroy}
```

Simply write:

```go
config := LifecycleConfig{}
config.SetCreateBeforeDestroy(true)
```

#### func (*LifecycleConfig) [SetPreventDestroy](/block_signature_extra.go#L77)

`func (c *LifecycleConfig) SetPreventDestroy(b bool)`

SetPreventDestroy is a simple helper to avoid having to create a boolean variable and then pass the pointer to it

E.g: instead of writing

```go
preventDestroy = false
config := LifecycleConfig{PreventDestroy: &preventDestroy}
```

Simply write:

```go
config := LifecycleConfig{}
config.SetPreventDestroy(true)
```

### type [ValueGenerator](/value_generator.go#L15)

`type ValueGenerator struct { ... }`

ValueGenerator is able to detect "ident" tokens and convert them into a special cty capsule.
Capsule will then be converted to `hclwrite.tokens`
It allows to write values like `var.my_var`, `locals.my_local` or `data.res_name.val_name` without any quotes.

```golang
basicStringValue := "basic_value"
localVal := "local.my_local"
varVal := "var.my_var"
dataVal := "data.my_data.my_property"
customStringValue := "custom.my_var"
identStringValue := "explicit_ident.foo"
identListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}

valGen := NewValueGenerator()
sig := NewEmptySignature("my_block")
sig.AppendAttribute("attr1", *valGen.ToString(&basicStringValue))
sig.AppendAttribute("attr2", *valGen.ToString(&localVal))
sig.AppendAttribute("attr3", *valGen.ToString(&varVal))
sig.AppendAttribute("attr4", *valGen.ToString(&dataVal))
sig.AppendAttribute("attr5", *valGen.ToString(&customStringValue))
sig.AppendEmptyLine()
sig.AppendAttribute("attr6", *valGen.ToIdent(&identStringValue))
sig.AppendAttribute("attr7", *valGen.ToIdentList(&identListStringValue))
customValGen := NewValueGenerator("custom.")
sig.AppendEmptyLine()
sig.AppendAttribute("attr8", *customValGen.ToString(&customStringValue))

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())
fmt.Println(string(hclFile.Bytes()))
```

 Output:

```
my_block {
  attr1 = "basic_value"
  attr2 = local.my_local
  attr3 = var.my_var
  attr4 = data.my_data.my_property
  attr5 = "custom.my_var"

  attr6 = explicit_ident.foo
  attr7 = [explicit_ident_item.foo, explicit_ident_item.bar]

  attr8 = custom.my_var
}
```

#### func [NewValueGenerator](/value_generator.go#L21)

`func NewValueGenerator(identPrefixList ...string) ValueGenerator`

NewValueGenerator returns a new ValueGenerator with the default 'ident' tokens matcher augmented with provided list
of token to consider as 'ident' tokens.

#### func [NewValueGeneratorWith](/value_generator.go#L26)

`func NewValueGeneratorWith(matcher IdentTokenMatcherInterface) ValueGenerator`

NewValueGeneratorWith returns a new ValueGenerator with the provided matcher.

#### func (*ValueGenerator) [FromString](/value_generator.go#L88)

`func (g *ValueGenerator) FromString(s *string, t cty.Type) *cty.Value`

FromString convert a string to `cty.Value` of the provided type
If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToBool](/value_generator.go#L56)

`func (g *ValueGenerator) ToBool(s *string) *cty.Value`

ToBool convert a string to `cty.Value` boolean which will be rendered as true or false value by terraform HCL
If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToIdent](/value_generator.go#L31)

`func (g *ValueGenerator) ToIdent(s *string) *cty.Value`

ToIdent converts a string to a special `cty.Value` capsule holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToIdentList](/value_generator.go#L40)

`func (g *ValueGenerator) ToIdentList(list *[]string) *cty.Value`

ToIdentList converts a list of string to `cty.Value` list containing capsules holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToNumber](/value_generator.go#L62)

`func (g *ValueGenerator) ToNumber(s *string) *cty.Value`

ToNumber convert a string to `cty.Value` number which will be rendered as numeric value by terraform HCL
If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToString](/value_generator.go#L50)

`func (g *ValueGenerator) ToString(s *string) *cty.Value`

ToString convert a string to `cty.Value` string which will be rendered as quoted string by terraform HCL
If the provided string is actually an 'ident' token, `cty.Value` will be a capsule holding `hclwrite.tokens`.

#### func (*ValueGenerator) [ToStringList](/value_generator.go#L68)

`func (g *ValueGenerator) ToStringList(list *[]string) *cty.Value`

ToStringList convert a string list to `cty.Value` string list which will be rendered as quoted string list by terraform HCL
If a provided string item is actually an 'ident' token, `cty.Value` item will be a capsule holding `hclwrite.tokens`.

## Sub Packages

* [testutils](./testutils)

* [tokens](./tokens): Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident)

## Examples

```golang
// Create a resource block
sig := NewEmptyResource("res_name", "res_id")
sig.AppendAttribute("attribute1", cty.StringVal("value1"))
sig.AppendEmptyLine()
sig.AppendAttribute("attribute2", cty.BoolVal(true))
sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
sig.AppendEmptyLine()

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "res_name" "res_id" {
  attribute1 = "value1"

  attribute2 = true
  attribute3 = -12.34

}
```

### Reorder

```golang
// Reorder an existing signature
sig := NewEmptyResource("res_name", "res_id")
sig.AppendAttribute("attribute1", cty.StringVal("value1"))
sig.AppendAttribute("attribute2", cty.BoolVal(true))
sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
sig.AppendEmptyLine()
sig.AppendAttribute("attribute4", cty.NumberIntVal(42))
sig.AppendAttribute("attribute5", cty.StringVal("value5"))

// ... Later in the code
// Re-order elements and remove empty lines
newElems := make(BodyElements, len(sig.GetElements()))
extraElemCont := 0
for _, v := range sig.GetElements() {
    if v.GetName() == "attribute1" {
        newElems[2] = v
    } else if v.GetName() == "attribute4" {
        newElems[1] = v
    } else if v.GetName() == "attribute3" {
        newElems[0] = v
    } else if !v.IsBodyEmptyLine() {
        newElems[3+extraElemCont] = v
        extraElemCont++
    }
}
sig.SetElements(newElems)

hclFile := hclwrite.NewEmptyFile()
hclFile.Body().AppendBlock(sig.Build())

fmt.Println(string(hclFile.Bytes()))
```

 Output:

```terraform
resource "res_name" "res_id" {
  attribute3 = -12.34
  attribute4 = 42
  attribute1 = "value1"
  attribute2 = true
  attribute5 = "value5"
}
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
