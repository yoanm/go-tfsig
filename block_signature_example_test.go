package tfsig_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

func Example() {
	// Create a resource block
	sig := tfsig.NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	sig.AppendEmptyLine()
	sig.AppendAttribute("attribute2", cty.BoolVal(true))
	sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
	sig.AppendEmptyLine()

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "res_name" "res_id" {
	//   attribute1 = "value1"
	//
	//   attribute2 = true
	//   attribute3 = -12.34
	//
	// }
}

func Example_enhance_existing() {
	// Enhance an existing signature
	sig := tfsig.NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	sig.AppendEmptyLine()
	sig.AppendAttribute("attribute2", cty.BoolVal(true))
	sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
	sig.AppendEmptyLine()

	// ... Later in the code
	// Enhance signature by removing empty lines and add an attribute below attribute2
	newElems := tfsig.BodyElements{}

	for _, elem := range sig.GetElements() {
		// Remove all empty lines
		if !elem.IsBodyEmptyLine() {
			newElems = append(newElems, elem)
		}
		// Append a new attribute right after attribute2
		if elem.GetName() == "attribute2" {
			newElems = append(newElems, tfsig.NewBodyAttribute("attribute22", cty.BoolVal(false)))
		}
	}

	sig.SetElements(newElems)

	// ... Finally
	hclFile := hclwrite.NewEmptyFile()

	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "res_name" "res_id" {
	//   attribute1  = "value1"
	//   attribute2  = true
	//   attribute22 = false
	//   attribute3  = -12.34
	// }
}

func Example_reorder() {
	// Reorder an existing signature
	sig := tfsig.NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	sig.AppendAttribute("attribute2", cty.BoolVal(true))
	sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
	sig.AppendEmptyLine()
	sig.AppendAttribute("attribute4", cty.NumberIntVal(42))
	sig.AppendAttribute("attribute5", cty.StringVal("value5"))

	// ... Later in the code
	// Re-order elements and remove empty lines
	newElems := make(tfsig.BodyElements, len(sig.GetElements()))
	extraElemCont := 0

	for _, elem := range sig.GetElements() {
		switch {
		case elem.GetName() == "attribute1":
			newElems[2] = elem
		case elem.GetName() == "attribute4":
			newElems[1] = elem
		case elem.GetName() == "attribute3":
			newElems[0] = elem
		case !elem.IsBodyEmptyLine():
			newElems[3+extraElemCont] = elem
			extraElemCont++
		}
	}

	sig.SetElements(newElems)

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "res_name" "res_id" {
	//   attribute3 = -12.34
	//   attribute4 = 42
	//   attribute1 = "value1"
	//   attribute2 = true
	//   attribute5 = "value5"
	// }
}
