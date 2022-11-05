package tfsig

import (
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/tfsig/testutils"
)

func TestBodyBlockBuild_panic(t *testing.T) {
	cases := map[string]struct {
		value BodyElement
	}{
		"AttributeBlock": {NewBodyAttribute("name", cty.StringVal("value"))},
		"BodyEmptyLine":  {NewBodyEmptyLine()},
	}

	expectedError := "Element is not a block"
	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				testutils.ExpectPanic(
					t,
					"Basic",
					func() {
						tc.value.Build()
					},
					expectedError,
				)
			},
		)
	}
}
