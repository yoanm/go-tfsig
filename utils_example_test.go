package tfsig_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

func ExampleAppendBlockIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	evenFn := func(i int) *hclwrite.Block {
		if i%2 == 0 {
			return nil
		}

		return tfsig.NewSignature(fmt.Sprintf("block%d", i)).Build()
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

		return tfsig.NewSignature(fmt.Sprintf("block%d", i)).Build()
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

func ExampleAppendAttributeIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	evenFn := func(i int) *cty.Value {
		if i%2 == 0 {
			return nil
		}

		val := cty.StringVal(fmt.Sprintf("val%d", i))

		return &val
	}

	sig := tfsig.NewResource("my_res", "res_id")

	tfsig.AppendAttributeIfNotNil(sig, "attr_0", evenFn(0))
	tfsig.AppendAttributeIfNotNil(sig, "attr_1", evenFn(1))
	tfsig.AppendAttributeIfNotNil(sig, "attr_2", evenFn(2))
	tfsig.AppendAttributeIfNotNil(sig, "attr_3", evenFn(3))
	tfsig.AppendAttributeIfNotNil(sig, "attr_4", evenFn(4))

	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "my_res" "res_id" {
	//   attr_1 = "val1"
	//   attr_3 = "val3"
	// }
}

func ExampleAppendChildIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	oddFn := func(i int) *tfsig.BlockSignature {
		if i%2 != 0 {
			return nil
		}

		return tfsig.NewSignature(fmt.Sprintf("block%d", i))
	}

	sig := tfsig.NewResource("my_res", "res_id")

	tfsig.AppendChildIfNotNil(sig, oddFn(0))
	tfsig.AppendChildIfNotNil(sig, oddFn(1))
	tfsig.AppendChildIfNotNil(sig, oddFn(2))
	tfsig.AppendChildIfNotNil(sig, oddFn(3))
	tfsig.AppendChildIfNotNil(sig, oddFn(4))

	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "my_res" "res_id" {
	//   block0 {
	//   }
	//
	//   block2 {
	//   }
	//
	//   block4 {
	//   }
	// }
}
