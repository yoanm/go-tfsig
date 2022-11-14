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
		tcase := tcase // For parallel execution

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
		tcase := tcase // For parallel execution

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

func TestToTerraformIdentifier(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		Value string
		Want  string
	}{
		"Already valid":          {"a_valid-identifier", "a_valid-identifier"},
		"Contains special chars": {"id_@&é\"'(§è!çà)^$`ù=:;,<#°*¨£%+/.?>|\\", "id_----------------------------------"},
		"Start with a number":    {"0-id", "_-id"},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				got := tfsig.ToTerraformIdentifier(tcase.Value)
				if got != tcase.Want {
					t.Errorf("wrong result for case %q: got %v, want %v", t.Name(), got, tcase.Want)
				}
			},
		)
	}
}
