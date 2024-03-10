package tfsig_test

import (
	"fmt"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
	"github.com/yoanm/go-tfsig/testutils"
)

func TestNewValueGenerator(t *testing.T) {
	t.Parallel()

	basicStringValue := "basic_value"
	customStringValue := "custom.my_var"
	localVal := "local.my_local"
	varVal := "var.my_var"
	dataVal := "data.my_data.my_property"
	boolVal := "true"
	intVal := "1"
	listVal := []string{basicStringValue, customStringValue, localVal, varVal, dataVal}

	cases := map[string]struct {
		value      tfsig.ValueGenerator
		goldenFile string
	}{
		"Default": {
			tfsig.NewValueGenerator(),
			"ValueGenerator.default",
		},
		"Enhanced": {
			tfsig.NewValueGenerator("custom."),
			"ValueGenerator.custom",
		},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				sig := tfsig.NewSignature("sig")
				sig.SetElements(
					tfsig.BodyElements{
						tfsig.NewBodyAttribute("basic", *tcase.value.ToString(&basicStringValue)),
						tfsig.NewBodyAttribute("custom", *tcase.value.ToString(&customStringValue)),
						tfsig.NewBodyAttribute("local", *tcase.value.ToString(&localVal)),
						tfsig.NewBodyAttribute("var", *tcase.value.ToString(&varVal)),
						tfsig.NewBodyAttribute("data", *tcase.value.ToString(&dataVal)),
						tfsig.NewBodyAttribute("bool", *tcase.value.ToBool(&boolVal)),
						tfsig.NewBodyAttribute("number", *tcase.value.ToNumber(&intVal)),
						tfsig.NewBodyAttribute("list_of_string", *tcase.value.ToStringList(&listVal)),
					},
				)

				if err := testutils.EnsureBlockFileEqualsGoldenFile(sig.Build(), tcase.goldenFile); err != nil {
					t.Errorf("Case \"%s\": %v", t.Name(), err)
				}
			},
		)
	}
}

func TestValueGenerator_nil(t *testing.T) {
	t.Parallel()

	valGen := tfsig.NewValueGenerator()
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

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				if actual := tcase.fn(); actual != nil {
					t.Errorf("wrong result for case %q: expected nil, got %v", t.Name(), actual)
				}
			},
		)
	}
}

func TestValueGenerator_panic(t *testing.T) {
	t.Parallel()

	valGen := tfsig.NewValueGenerator()
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
	t.Parallel()

	valGen := tfsig.NewValueGenerator()
	if actual := valGen.ToIdent(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}

func TestToIdentList_nil(t *testing.T) {
	t.Parallel()

	valGen := tfsig.NewValueGenerator()
	if actual := valGen.ToIdentList(nil); actual != nil {
		t.Errorf("wrong result: expected nil, got %v", actual)
	}
}
