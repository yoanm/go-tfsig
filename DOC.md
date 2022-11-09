# tfsig

Package tfsig is a wrapper for Terraform HCL language (hclwrite).
It provides ability to generate block signature which are way easier to manipulate and alter than hclwrite.tokens type

## Sub Packages

* [testutils](./testutils)

* [tokens](./tokens): Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident) It also provides an easy way to encapsulate hclwrite tokens into a cty.Value and a function (Generate()) to manage those type of value

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

```
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

```
resource "res_name" "res_id" {
  attribute3 = -12.34
  attribute4 = 42
  attribute1 = "value1"
  attribute2 = true
  attribute5 = "value5"
}
```

### DependsOn

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

```
resource "res_name" "res_id" {
  attribute1 = "value1"

  depends_on = [another_res.res_id, another_another_res.res_id]
}
```

### Lifecycle

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

```
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

### AppendBlockIfNotNil

AppendBlockIfNotNil appends the provided block to the provided body only if block is not nil
It simply avoids an if in your code

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

### AppendNewLineAndBlockIfNotNil

AppendNewLineAndBlockIfNotNil appends an empty line followed by provided block to the provided body only if block is not nil
It simply avoids an if in your code

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

### ToTerraformIdentifier

ToTerraformIdentifier converts a string to a terraform identifier, by converting not allowed characters to '-'
And if provided value starts with a character not allowed as first character, it replaces it by '_'

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

### ValueGenerator

ValueGenerator is able to detect "ident" tokens and convert them into a special cty capsule.
Capsule will then be converted to hclwrite.tokens
It allows to write values like 'var.my_var', 'locals.my_local' or 'data.res_name.val_name' without any quotes

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

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
