package tfsig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zclconf/go-cty/cty"

	"github.com/yoanm/go-tfsig/testutils"
)

func TestBuild_panic(t *testing.T) {
	cases := map[string]struct {
		value BodyElement
	}{
		"AttributeBlock": {NewBodyAttribute("name", cty.StringVal("value"))},
		"BodyEmptyLine":  {NewBodyEmptyLine()},
	}

	expectedError := "element is not a body block"
	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				testutils.ExpectPanic(
					t,
					tcname,
					func() {
						tc.value.Build()
					},
					expectedError,
				)
			},
		)
	}
}

func TestGetBodyAttribute(t *testing.T) {
	attr := cty.StringVal("value")
	elem := NewBodyAttribute("name", attr)
	if !reflect.DeepEqual(attr, *elem.GetBodyAttribute()) {
		t.Errorf("Mismatch want %#v, got %#v", attr, *elem.GetBodyAttribute())
	}
}

func TestGetBodyAttribute_panic(t *testing.T) {
	cases := map[string]struct {
		value BodyElement
	}{
		"BodyBlock":     {NewBodyBlock(NewEmptyResource("res", "id"))},
		"BodyEmptyLine": {NewBodyEmptyLine()},
	}

	expectedError := "element is not a body attribute"
	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				testutils.ExpectPanic(
					t,
					tcname,
					func() {
						tc.value.GetBodyAttribute()
					},
					expectedError,
				)
			},
		)
	}
}

func TestGetBodyBlock(t *testing.T) {
	block := NewEmptyResource("res", "id")
	elem := NewBodyBlock(block)

	if !reflect.DeepEqual(block, elem.GetBodyBlock()) {
		t.Errorf("Mismatch want %#v, got %#v", block, elem.GetBodyBlock())
	} else if fmt.Sprintf("%p", block) != fmt.Sprintf("%p", elem.GetBodyBlock()) {
		t.Errorf("Mismatch want pointer to %p, got %p", block, elem.GetBodyBlock())
	}
}

func TestGetBodyBlock_panic(t *testing.T) {
	cases := map[string]struct {
		value BodyElement
	}{
		"AttributeBlock": {NewBodyAttribute("name", cty.StringVal("value"))},
		"BodyEmptyLine":  {NewBodyEmptyLine()},
	}

	expectedError := "element is not a body block"
	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				testutils.ExpectPanic(
					t,
					tcname,
					func() {
						tc.value.GetBodyBlock()
					},
					expectedError,
				)
			},
		)
	}
}
