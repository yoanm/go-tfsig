package tfsig

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/yoanm/go-tfsig/testutils"
)

func TestAppendNewLineAndBlockIfNotNil(t *testing.T) {
	testBlock := hclwrite.NewBlock("source", nil)
	cases := map[string]struct {
		Value *hclwrite.Block
		Want  string
	}{
		"Nil block":     {nil, ""},
		"Not nil block": {testBlock, "\nsource {\n}\n"},
	}

	for tcname, tc := range cases {
		file := hclwrite.NewEmptyFile()
		AppendNewLineAndBlockIfNotNil(file.Body(), tc.Value)
		if err := testutils.EnsureFileContentEquals(file, tc.Want); err != nil {
			t.Errorf("Case \"%s\": %v", tcname, err)
		}
	}
}

func TestAppendBlockIfNotNil(t *testing.T) {
	testBlock := hclwrite.NewBlock("source", nil)
	cases := map[string]struct {
		Value *hclwrite.Block
		Want  string
	}{
		"Nil block":     {nil, ""},
		"Not nil block": {testBlock, "source {\n}\n"},
	}

	for tcname, tc := range cases {
		file := hclwrite.NewEmptyFile()
		AppendBlockIfNotNil(file.Body(), tc.Value)
		if err := testutils.EnsureFileContentEquals(file, tc.Want); err != nil {
			t.Errorf("Case \"%s\": %v", tcname, err)
		}
	}
}

func TestToTerraformIdentifier(t *testing.T) {
	cases := map[string]struct {
		Value string
		Want  string
	}{
		"Already valid":          {"a_valid-identifier", "a_valid-identifier"},
		"Contains special chars": {"id_@&é\"'(§è!çà)^$`ù=:;,<#°*¨£%+/.?>|\\", "id_----------------------------------"},
		"Start with a number":    {"0-id", "_-id"},
	}

	for caseName, tc := range cases {
		got := ToTerraformIdentifier(tc.Value)
		if got != tc.Want {
			t.Errorf("wrong result for case %q: got %v, want %v", caseName, got, tc.Want)
		}
	}
}
