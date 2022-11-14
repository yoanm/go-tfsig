package tfsig_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig"
	"github.com/yoanm/go-tfsig/testutils"
)

func TestBuild_panic(t *testing.T) {
	t.Parallel()

	expectedError := "element is not a body block"
	cases := map[string]struct {
		value tfsig.BodyElement
	}{
		"AttributeBlock": {tfsig.NewBodyAttribute("name", cty.StringVal("value"))},
		"BodyEmptyLine":  {tfsig.NewBodyEmptyLine()},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				testutils.ExpectPanic(
					t,
					t.Name(),
					func() {
						tcase.value.Build()
					},
					expectedError,
				)
			},
		)
	}
}

func TestGetBodyAttribute(t *testing.T) {
	t.Parallel()

	attr := cty.StringVal("value")
	elem := tfsig.NewBodyAttribute("name", attr)

	if !reflect.DeepEqual(attr, *elem.GetBodyAttribute()) {
		t.Errorf("Mismatch want %#v, got %#v", attr, *elem.GetBodyAttribute())
	}
}

func TestGetBodyAttribute_panic(t *testing.T) {
	t.Parallel()

	expectedError := "element is not a body attribute"
	cases := map[string]struct {
		value tfsig.BodyElement
	}{
		"BodyBlock":     {tfsig.NewBodyBlock(tfsig.NewEmptyResource("res", "id"))},
		"BodyEmptyLine": {tfsig.NewBodyEmptyLine()},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				testutils.ExpectPanic(
					t,
					t.Name(),
					func() {
						tcase.value.GetBodyAttribute()
					},
					expectedError,
				)
			},
		)
	}
}

func TestGetBodyBlock(t *testing.T) {
	t.Parallel()

	block := tfsig.NewEmptyResource("res", "id")
	elem := tfsig.NewBodyBlock(block)

	if !reflect.DeepEqual(block, elem.GetBodyBlock()) {
		t.Errorf("Mismatch want %#v, got %#v", block, elem.GetBodyBlock())
	} else if fmt.Sprintf("%p", block) != fmt.Sprintf("%p", elem.GetBodyBlock()) {
		t.Errorf("Mismatch want pointer to %p, got %p", block, elem.GetBodyBlock())
	}
}

func TestGetBodyBlock_panic(t *testing.T) {
	t.Parallel()

	expectedError := "element is not a body block"
	cases := map[string]struct {
		value tfsig.BodyElement
	}{
		"AttributeBlock": {tfsig.NewBodyAttribute("name", cty.StringVal("value"))},
		"BodyEmptyLine":  {tfsig.NewBodyEmptyLine()},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				testutils.ExpectPanic(
					t,
					t.Name(),
					func() {
						tcase.value.GetBodyBlock()
					},
					expectedError,
				)
			},
		)
	}
}
