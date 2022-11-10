package tokens

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/testutils"
)

func TestGenerate_nil(t *testing.T) {
	got := Generate(nil)
	hclFile := hclwrite.NewEmptyFile()
	hclFile.Body().AppendUnstructuredTokens(got)
	actual := string(hclFile.Bytes())
	if actual != "" {
		t.Errorf("expected empty string, got  %s", actual)
	}
}

func TestGenerateFromIterable_panic(t *testing.T) {
	expectedError := "expected a collection type but got cty.String"
	testutils.ExpectPanic(
		t,
		"Basic",
		func() {
			GenerateFromIterable(nil, cty.String)
		},
		expectedError,
	)
}

func TestSplitIterable_panic(t *testing.T) {
	expectedError := "expected an iterable type but got cty.String"
	testutils.ExpectPanic(
		t,
		"Basic",
		func() {
			SplitIterable(cty.StringVal(""))
		},
		expectedError,
	)
}
