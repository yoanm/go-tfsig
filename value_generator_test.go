package tfsig

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/testutils"
)

func TestNewValueGenerator(t *testing.T) {
	basicStringValue := "basic_value"
	customStringValue := "custom.my_var"
	localVal := "local.my_local"
	varVal := "var.my_var"
	dataVal := "data.my_data.my_property"
	boolVal := "true"
	intVal := "1"
	listVal := []string{basicStringValue, customStringValue, localVal, varVal, dataVal}

	cases := map[string]struct {
		value      ValueGenerator
		goldenFile string
	}{
		"Default": {
			NewValueGenerator(),
			"ValueGenerator.default",
		},
		"Enhanced": {
			NewValueGenerator("custom."),
			"ValueGenerator.custom",
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				sig := NewSignature(
					"sig",
					nil,
					BodyElements{
						NewBodyAttribute("basic", *tc.value.ToString(&basicStringValue)),
						NewBodyAttribute("custom", *tc.value.ToString(&customStringValue)),
						NewBodyAttribute("local", *tc.value.ToString(&localVal)),
						NewBodyAttribute("var", *tc.value.ToString(&varVal)),
						NewBodyAttribute("data", *tc.value.ToString(&dataVal)),
						NewBodyAttribute("bool", *tc.value.ToBool(&boolVal)),
						NewBodyAttribute("number", *tc.value.ToNumber(&intVal)),
						NewBodyAttribute("list_of_string", *tc.value.ToStringList(&listVal)),
					},
				)
				if err := testutils.EnsureBlockFileEqualsGoldenFile(sig.Build(), tc.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", tcname, err)
				}
			},
		)
	}
}

func TestValueGenerator_nil(t *testing.T) {
	valGen := NewValueGenerator()
	cases := map[string]struct {
		fn func() *cty.Value
	}{
		"String": {
			func() *cty.Value {
				return valGen.ToString(nil)
			},
		},
		"Bool": {
			func() *cty.Value {
				return valGen.ToBool(nil)
			},
		},
		"Number": {
			func() *cty.Value {
				return valGen.ToNumber(nil)
			},
		},
		"List of string": {
			func() *cty.Value {
				return valGen.ToStringList(nil)
			},
		},
		"FromString": {
			func() *cty.Value {
				return valGen.FromString(nil, cty.List(cty.Bool))
			},
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				if actual := tc.fn(); actual != nil {
					t.Errorf("wrong result for case %q: expected nil, got %v", tcname, actual)
				}
			},
		)
	}
}

func TestValueGenerator_panic(t *testing.T) {
	valGen := NewValueGenerator()
	val := "a_value"
	ctyType := cty.Map(cty.String)
	expectedError := fmt.Sprintf("Unable to convert \"%s\" to a %s", val, ctyType.FriendlyName())

	testutils.ExpectPanic(
		t,
		"Basic",
		func() {
			valGen.FromString(&val, ctyType)
		},
		expectedError,
	)
}

func TestToIdent_nil(t *testing.T) {
	valGen := NewValueGenerator()
	if actual := valGen.ToIdent(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}

func TestToIdentList_nil(t *testing.T) {
	valGen := NewValueGenerator()
	if actual := valGen.ToIdentList(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}
