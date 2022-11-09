package tokens

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func ExampleNewIdentValue() {
	identStringValue := "explicit_ident.foo"
	value := NewIdentValue(identStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", Generate(value))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// attr = explicit_ident.foo
}

func ExampleNewIdentListValue() {
	identListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}
	value := NewIdentListValue(identListStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", Generate(value))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// attr = [explicit_ident_item.foo, explicit_ident_item.bar]
}
