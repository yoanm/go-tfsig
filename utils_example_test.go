package tfsig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func ExampleToTerraformIdentifier() {
	fmt.Printf("a_valid-identifier becomes %s\n", ToTerraformIdentifier("a_valid-identifier"))
	fmt.Printf(".github becomes %s\n", ToTerraformIdentifier(".github"))
	fmt.Printf("an.identifier becomes %s\n", ToTerraformIdentifier("an.identifier"))
	fmt.Printf("0id becomes %s\n", ToTerraformIdentifier("0id"))

	// Output:
	// a_valid-identifier becomes a_valid-identifier
	// .github becomes _github
	// an.identifier becomes an-identifier
	// 0id becomes _id
}

func ExampleAppendBlockIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	evenFn := func(i int) *hclwrite.Block {
		if i%2 == 0 {
			return nil
		} else {
			return NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
		}
	}

	AppendBlockIfNotNil(hclFile.Body(), evenFn(0))
	AppendBlockIfNotNil(hclFile.Body(), evenFn(1))
	AppendBlockIfNotNil(hclFile.Body(), evenFn(2))
	AppendBlockIfNotNil(hclFile.Body(), evenFn(3))
	AppendBlockIfNotNil(hclFile.Body(), evenFn(4))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// block1 {
	// }
	// block3 {
	// }
}

func ExampleAppendNewLineAndBlockIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	oddFn := func(i int) *hclwrite.Block {
		if i%2 != 0 {
			return nil
		} else {
			return NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
		}
	}

	AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(0))
	AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(1))
	AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(2))
	AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(3))
	AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(4))

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// block0 {
	// }
	//
	// block2 {
	// }
	//
	// block4 {
	// }
}
