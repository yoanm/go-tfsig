package tfsig

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/yoanm/tfsig/testutils"
	"github.com/yoanm/tfsig/tokens"
)

func TestNewEmptyResource(t *testing.T) {
	resFull := NewEmptyResource("res_name", "res_id")
	resFull.AppendAttribute("attribute1", cty.StringVal("value1"))
	resFull.AppendEmptyLine()

	block1 := NewEmptySignature("block1", "block1_label1", "block1_label2", "block1_label3", "block1_label4")
	block1.AppendAttribute("attribute21", cty.BoolVal(true))
	block1.AppendAttribute("attribute22", cty.NumberIntVal(3))

	block11 := NewEmptySignature("block11")
	listAttr, _ := gocty.ToCtyValue([]string{"A", "B"}, cty.List(cty.String))
	block11.AppendAttribute("attribute211", listAttr)
	block11.AppendEmptyLine()
	block11.AppendAttribute("attribute212", cty.NumberIntVal(-10))
	block11.AppendEmptyLine()

	block1.AppendChild(block11)

	resFull.AppendChild(block1)
	resFull.AppendEmptyLine()
	resFull.AppendAttribute("attribute2", cty.StringVal("value2"))

	cases := map[string]struct {
		value      *BlockSignature
		goldenFile string
	}{
		"Full": {
			resFull,
			"resource.full",
		},
		"Empty": {
			NewEmptyResource("res_name", "res_id"),
			"resource.empty",
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				if err := testutils.EnsureBlockFileEqualsGoldenFile(tc.value.Build(), tc.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", tcname, err)
				}
			},
		)
	}
}

func TestBlockSignature_BuildTokens(t *testing.T) {
	res := NewEmptyResource("res_name", "res_id")
	res.AppendAttribute("attribute1", cty.StringVal("value1"))
	res.AppendEmptyLine()

	block1 := NewEmptySignature("block1", "block1_label1", "block1_label2", "block1_label3", "block1_label4")
	block1.AppendAttribute("attribute21", cty.BoolVal(true))
	block1.AppendAttribute("attribute22", cty.NumberIntVal(3))

	block11 := NewEmptySignature("block11")
	listAttr, _ := gocty.ToCtyValue([]string{"A", "B"}, cty.List(cty.String))
	block11.AppendAttribute("attribute211", listAttr)
	block11.AppendEmptyLine()
	block11.AppendAttribute("attribute212", cty.NumberIntVal(-10))
	block11.AppendEmptyLine()

	block1.AppendChild(block11)

	res.AppendChild(block1)
	res.AppendEmptyLine()
	res.AppendAttribute("attribute2", cty.StringVal("value2"))

	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendUnstructuredTokens(res.BuildTokens())

	// BuildTokens remove trailing new line char but original file has one
	// => Re add it
	hclFile.Body().AppendUnstructuredTokens(tokens.NewLineTokens())

	if err := testutils.EnsureFileEqualsGoldenFile(hclFile, "resource.full"); err != nil {
		t.Error(err)
	}
}
