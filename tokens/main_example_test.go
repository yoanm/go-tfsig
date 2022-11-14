package tokens_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-tfsig/tokens"
)

func ExampleNewIdentValue() {
	identStringValue := "explicit_ident.foo"
	value := tokens.NewIdentValue(identStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", tokens.Generate(value))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// attr = explicit_ident.foo
}

func ExampleNewIdentListValue() {
	identListStringValue := []string{"explicit_ident_item.foo", "explicit_ident_item.bar"}
	value := tokens.NewIdentListValue(identListStringValue)

	// ... Later in the code
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().SetAttributeRaw("attr", tokens.Generate(value))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// attr = [explicit_ident_item.foo, explicit_ident_item.bar]
}
