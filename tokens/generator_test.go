package tokens

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/yoanm/tfsig/testutils"
)

func TestGenerate(t *testing.T) {
	listOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*NewIdentValue("value1"),
			*NewIdentValue("value2"),
		},
		cty.List(hclwriteTokensCtyType),
	)
	if err != nil {
		t.Fatal(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*NewIdentValue("value1"),
			*NewIdentValue("value2"),
		},
		cty.Set(hclwriteTokensCtyType),
	)
	if err != nil {
		t.Fatal(err)
	}

	objectWithCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *NewIdentValue("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Object(
			map[string]cty.Type{
				"A": hclwriteTokensCtyType,
				"B": cty.String,
			},
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *NewIdentValue("A_value"),
			"B": *NewIdentValue("B_value"),
		},
		cty.Map(hclwriteTokensCtyType),
	)
	if err != nil {
		t.Fatal(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			*NewIdentValue("B_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, hclwriteTokensCtyType, cty.Number}),
	)
	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		Value cty.Value
		Want  string
	}{
		"Null": {
			cty.NilVal,
			"null",
		},
		"Ident": {
			*NewIdentValue("TeSt"),
			"TeSt",
		},
		"String": {
			cty.StringVal("TeSt"),
			"\"TeSt\"",
		},
		"Positive number": {
			cty.NumberIntVal(12),
			"12",
		},
		"Negative number": {
			cty.NumberFloatVal(-12.23),
			"-12.23",
		},
		"Boolean": {
			cty.BoolVal(false),
			"false",
		},
		"List of capsule": {
			listOfCapsule,
			"[value1, value2]",
		},
		"Set of capsule": {
			setOfCapsule,
			"[value1, value2]",
		},
		"Object with capsule": {
			objectWithCapsule,
			"{\n  \"A\" = A_value\n  \"B\" = \"B_value\"\n}",
		},
		"Map of capsule": {
			mapOfCapsule,
			"{\n  \"A\" = A_value\n  \"B\" = B_value\n}",
		},
		"Tuple with capsule": {
			tupleWithCapsule,
			"[\"A_value\", B_value, 2]",
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				got := Generate(&tc.Value)
				hclFile := hclwrite.NewEmptyFile()
				hclFile.Body().AppendUnstructuredTokens(got)
				if validateErr := testutils.EnsureFileContentEquals(hclFile, tc.Want); validateErr != nil {
					t.Errorf("Case \"%s\": %v", tcname, validateErr)
				}
			},
		)
	}
}

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
			GenerateFromIterable(cty.String, nil)
		},
		expectedError,
	)
}

func TestSplitIterable(t *testing.T) {
	list, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.List(cty.String),
	)
	if err != nil {
		t.Fatal(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.Set(cty.String),
	)
	if err != nil {
		t.Fatal(err)
	}

	objectWithCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": cty.NumberIntVal(2),
			"B": cty.StringVal("B_value"),
		},
		cty.Object(
			map[string]cty.Type{
				"A": cty.Number,
				"B": cty.String,
			},
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": cty.StringVal("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Map(cty.String),
	)
	if err != nil {
		t.Fatal(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, cty.Number}),
	)
	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		Value        cty.Value
		WantStart    string
		WantElements string
		WantEnd      string
	}{
		"List": {
			list,
			"[",
			"\"value1\", \"value2\"",
			"]",
		},
		"Set": {
			setOfCapsule,
			"[",
			"\"value1\", \"value2\"",
			"]",
		},
		"Object": {
			objectWithCapsule,
			"{",
			"\nA = 2\nB = \"B_value\"\n",
			"}",
		},
		"Map": {
			mapOfCapsule,
			"{",
			"\nA = \"A_value\"\nB = \"B_value\"\n",
			"}",
		},
		"Tuple": {
			tupleWithCapsule,
			"[",
			"\"A_value\", 2",
			"]",
		},
	}

	for tcname, tc := range cases {
		t.Run(
			tcname,
			func(t *testing.T) {
				start, elements, end := SplitIterable(tc.Value)
				hclFile := hclwrite.NewEmptyFile()
				hclFile.Body().AppendUnstructuredTokens(start)
				if validateErr := testutils.EnsureFileContentEquals(hclFile, tc.WantStart); validateErr != nil {
					t.Errorf("Case \"%s\": Start:%v", tcname, validateErr)
				}
				hclFile.Body().Clear()
				hclFile.Body().AppendUnstructuredTokens(elements)
				if validateErr := testutils.EnsureFileContentEquals(hclFile, tc.WantElements); validateErr != nil {
					t.Errorf("Case \"%s\": Elements:%v", tcname, validateErr)
				}
				hclFile.Body().Clear()
				hclFile.Body().AppendUnstructuredTokens(end)
				if validateErr := testutils.EnsureFileContentEquals(hclFile, tc.WantEnd); validateErr != nil {
					t.Errorf("Case \"%s\": End:%v", tcname, validateErr)
				}
			},
		)
	}
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
