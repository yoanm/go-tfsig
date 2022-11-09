# tokens

Package tokens provides an easy way to create common hclwrite tokens (such as new line, comma, equal sign, ident)
It also provides an easy way to encapsulate hclwrite tokens into a cty.Value and a function (Generate()) to manage those type of value

## Examples

### NewIdentListValue

NewIdentListValue tales a list of string which should be all considered as 'ident' tokens
and converts them into a cty list containing special cty.Value capsule

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

### NewIdentValue

NewIdentValue takes a string which should be considered as 'ident' token and converts it to a special cty.Value capsule

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

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
