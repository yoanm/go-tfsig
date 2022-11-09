package tokens

import (
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/testutils"
)

func TestFromValue_panic(t *testing.T) {
	expectedError := "error during conversion from cty.Value to hclwrite.Tokens: list or set value is required"
	testutils.ExpectPanic(
		t,
		"Basic",
		func() {
			FromValue(cty.StringVal("A"))
		},
		expectedError,
	)
}

func TestNewIdentListValue_nil(t *testing.T) {
	if actual := NewIdentListValue(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}
