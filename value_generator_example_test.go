package tfsig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func ExampleValueGenerator() {
	basicStringValue := "basic_value"
	localVal := "local.my_local"
	varVal := "var.my_var"
	dataVal := "data.my_data.my_property"
	customStringValue := "custom.my_var"

	valGen := NewValueGenerator()
	sig := NewEmptySignature("my_block")
	sig.AppendAttribute("attr1", *valGen.ToString(&basicStringValue))
	sig.AppendAttribute("attr2", *valGen.ToString(&localVal))
	sig.AppendAttribute("attr3", *valGen.ToString(&varVal))
	sig.AppendAttribute("attr4", *valGen.ToString(&dataVal))
	sig.AppendAttribute("attr5", *valGen.ToString(&customStringValue))
	customValGen := NewValueGenerator("custom.")
	sig.AppendEmptyLine()
	sig.AppendAttribute("attr6", *customValGen.ToString(&customStringValue))

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(sig.Build())
	fmt.Println(string(hclFile.Bytes()))

	// Output:
	// my_block {
	//   attr1 = "basic_value"
	//   attr2 = local.my_local
	//   attr3 = var.my_var
	//   attr4 = data.my_data.my_property
	//   attr5 = "custom.my_var"
	//
	//   attr6 = custom.my_var
	// }
}
