package tfsig_test

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
)

func ExampleBlockSignature_appendAttributeIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	evenFn := func(i int) *cty.Value {
		if i%2 == 0 {
			return nil
		}

		val := cty.StringVal(fmt.Sprintf("val%d", i))

		return &val
	}

	sig := tfsig.NewEmptyResource("my_res", "res_id")

	sig.AppendAttributeIfNotNil("attr_0", evenFn(0))
	sig.AppendAttributeIfNotNil("attr_1", evenFn(1))
	sig.AppendAttributeIfNotNil("attr_2", evenFn(2))
	sig.AppendAttributeIfNotNil("attr_3", evenFn(3))
	sig.AppendAttributeIfNotNil("attr_4", evenFn(4))

	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "my_res" "res_id" {
	//   attr_1 = "val1"
	//   attr_3 = "val3"
	// }
}

func ExampleBlockSignature_appendChildIfNotNil() {
	hclFile := hclwrite.NewEmptyFile()
	oddFn := func(i int) *tfsig.BlockSignature {
		if i%2 != 0 {
			return nil
		}

		return tfsig.NewEmptySignature(fmt.Sprintf("block%d", i))
	}

	sig := tfsig.NewEmptyResource("my_res", "res_id")

	sig.AppendChildIfNotNil(oddFn(0))
	sig.AppendChildIfNotNil(oddFn(1))
	sig.AppendChildIfNotNil(oddFn(2))
	sig.AppendChildIfNotNil(oddFn(3))
	sig.AppendChildIfNotNil(oddFn(4))

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
