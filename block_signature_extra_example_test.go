package tfsig

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func ExampleBlockSignature_DependsOn() {
	// resource with 'depends_on' directive
	sig := NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	sig.DependsOn([]string{"another_res.res_id", "another_another_res.res_id"})

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(sig.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "res_name" "res_id" {
	//   attribute1 = "value1"
	//
	//   depends_on = [another_res.res_id, another_another_res.res_id]
	// }
}

func ExampleBlockSignature_Lifecycle() {
	// resource with 'lifecycle' directive
	sig := NewEmptyResource("res_name", "res_id")
	sig.AppendAttribute("attribute1", cty.StringVal("value1"))
	config := LifecycleConfig{}
	config.SetCreateBeforeDestroy(true)
	config.SetPreventDestroy(false)
	sig.Lifecycle(config)

	sig2 := NewEmptyResource("res2_name", "res2_id")
	sig2.AppendAttribute("attribute1", cty.StringVal("value1"))
	config2 := LifecycleConfig{
		IgnoreChanges: []string{"attribute1"},
		Postcondition: &LifecycleCondition{
			condition:    "res_name.res_id.attribute1 != \"value1\"",
			errorMessage: "res_name.res_id.attribute1 must equal \"value1\"",
		},
	}
	sig2.Lifecycle(config2)

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendBlock(sig.Build())
	hclFile.Body().AppendBlock(sig2.Build())

	fmt.Println(string(hclFile.Bytes()))
	// Output:
	// resource "res_name" "res_id" {
	//   attribute1 = "value1"
	//
	//   lifecycle {
	//     create_before_destroy = true
	//     prevent_destroy       = false
	//   }
	// }
	// resource "res2_name" "res2_id" {
	//   attribute1 = "value1"
	//
	//   lifecycle {
	//     ignore_changes = [attribute1]
	//     postcondition {
	//       condition     = res_name.res_id.attribute1 != "value1"
	//       error_message = "res_name.res_id.attribute1 must equal \"value1\""
	//     }
	//   }
	// }
}
