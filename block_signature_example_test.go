package tfsig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func Example() {
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
	sig := NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	sig.AppendEmptyLine()
	sig.AppendAttribute("attribute2", cty.BoolVal(true))
	sig.AppendAttribute("attribute3", cty.NumberFloatVal(-12.34))
	sig.AppendEmptyLine()

	// ... Later in the code
	// Enhance signature by removing empty lines and add an attribute below attribute2
	newElems := BodyElements{}
	for _, v := range sig.GetElements() {
		// Remove all empty lines
		if !v.IsBodyEmptyLine() {
			newElems = append(newElems, v)
		}
		// Append a new attribute right after attribute2
		if v.GetName() == "attribute2" {
			newElems = append(newElems, NewBodyAttribute("attribute22", cty.BoolVal(false)))
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
	// Output:
	// resource "res_name" "res_id" {
	//   attribute3 = -12.34
	//   attribute4 = 42
	//   attribute1 = "value1"
	//   attribute2 = true
	//   attribute5 = "value5"
	// }
}
