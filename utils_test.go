package tfsig_test

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-tfsig"
	"github.com/yoanm/go-tfsig/testutils"
)

func TestAppendNewLineAndBlockIfNotNil(t *testing.T) {
	t.Parallel()

	testBlock := hclwrite.NewBlock("source", nil)
	cases := map[string]struct {
		Value *hclwrite.Block
		Want  string
	}{
		"Nil block":     {nil, ""},
		"Not nil block": {testBlock, "\nsource {\n}\n"},
	}

	for tcname, tcase := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				file := hclwrite.NewEmptyFile()
				tfsig.AppendNewLineAndBlockIfNotNil(file.Body(), tcase.Value)

				if err := testutils.EnsureFileContentEquals(file, tcase.Want); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

func TestAppendBlockIfNotNil(t *testing.T) {
	t.Parallel()

	testBlock := hclwrite.NewBlock("source", nil)
	cases := map[string]struct {
		Value *hclwrite.Block
		Want  string
	}{
		"Nil block":     {nil, ""},
		"Not nil block": {testBlock, "source {\n}\n"},
	}

	for tcname, tcase := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				file := hclwrite.NewEmptyFile()

				tfsig.AppendBlockIfNotNil(file.Body(), tcase.Value)

				if err := testutils.EnsureFileContentEquals(file, tcase.Want); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}
