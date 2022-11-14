package tokens_test

import (
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/testutils"
	"github.com/yoanm/go-tfsig/tokens"
)

func TestFromValue_panic(t *testing.T) {
	t.Parallel()

	expectedError := "error during conversion from cty.Value to hclwrite.Tokens: list or set value is required"
	testutils.ExpectPanic(
		t,
		"Basic",
		func() {
			tokens.FromValue(cty.StringVal("A"))
		},
		expectedError,
	)
}

func TestNewIdentListValue_nil(t *testing.T) {
	t.Parallel()

	if actual := tokens.NewIdentListValue(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}
