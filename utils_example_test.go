package tfsig_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-tfsig"
)

func ExampleToTerraformIdentifier() {
	fmt.Printf("a_valid-identifier becomes %s\n", tfsig.ToTerraformIdentifier("a_valid-identifier"))
	fmt.Printf(".github becomes %s\n", tfsig.ToTerraformIdentifier(".github"))
	fmt.Printf("an.identifier becomes %s\n", tfsig.ToTerraformIdentifier("an.identifier"))
	fmt.Printf("0id becomes %s\n", tfsig.ToTerraformIdentifier("0id"))

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
		}

		return tfsig.NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
	}

	tfsig.AppendBlockIfNotNil(hclFile.Body(), evenFn(0))
	tfsig.AppendBlockIfNotNil(hclFile.Body(), evenFn(1))
	tfsig.AppendBlockIfNotNil(hclFile.Body(), evenFn(2))
	tfsig.AppendBlockIfNotNil(hclFile.Body(), evenFn(3))
	tfsig.AppendBlockIfNotNil(hclFile.Body(), evenFn(4))

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
		}

		return tfsig.NewEmptySignature(fmt.Sprintf("block%d", i)).Build()
	}

	tfsig.AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(0))
	tfsig.AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(1))
	tfsig.AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(2))
	tfsig.AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(3))
	tfsig.AppendNewLineAndBlockIfNotNil(hclFile.Body(), oddFn(4))

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
