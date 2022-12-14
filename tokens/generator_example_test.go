package tokens_test

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/yoanm/go-tfsig/tokens"
)

func ExampleGenerate() {
	listOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*tokens.NewIdentValue("value1"),
			*tokens.NewIdentValue("value2"),
		},
		cty.List(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			*tokens.NewIdentValue("value1"),
			*tokens.NewIdentValue("value2"),
		},
		cty.Set(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	objectWithCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *tokens.NewIdentValue("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Object(
			map[string]cty.Type{
				"A": cty.DynamicPseudoType,
				"B": cty.String,
			},
		),
	)
	if err != nil {
		panic(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": *tokens.NewIdentValue("A_value"),
			"B": *tokens.NewIdentValue("B_value"),
		},
		cty.Map(cty.DynamicPseudoType),
	)
	if err != nil {
		panic(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			*tokens.NewIdentValue("B_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, cty.DynamicPseudoType, cty.Number}),
	)
	if err != nil {
		panic(err)
	}

	stringVal := cty.StringVal("TeSt")
	numberIntVal := cty.NumberIntVal(12)
	numberFloatVal := cty.NumberFloatVal(-12.23)
	boolVal := cty.BoolVal(false)

	fmt.Printf("Null: %#v\n", string(tokens.Generate(&cty.NilVal).Bytes()))
	fmt.Printf("Ident: %#v\n", string(tokens.Generate(tokens.NewIdentValue("TeSt")).Bytes()))
	fmt.Printf("String: %#v\n", string(tokens.Generate(&stringVal).Bytes()))
	fmt.Printf("Positive number: %#v\n", string(tokens.Generate(&numberIntVal).Bytes()))
	fmt.Printf("Negative number: %#v\n", string(tokens.Generate(&numberFloatVal).Bytes()))
	fmt.Printf("Boolean: %#v\n", string(tokens.Generate(&boolVal).Bytes()))
	fmt.Printf("List of capsule: %#v\n", string(tokens.Generate(&listOfCapsule).Bytes()))
	fmt.Printf("Set of capsule: %#v\n", string(tokens.Generate(&setOfCapsule).Bytes()))
	fmt.Printf("Object with capsule: %#v\n", string(tokens.Generate(&objectWithCapsule).Bytes()))
	fmt.Printf("Map of capsule: %#v\n", string(tokens.Generate(&mapOfCapsule).Bytes()))
	fmt.Printf("Tuple with capsule: %#v\n", string(tokens.Generate(&tupleWithCapsule).Bytes()))

	// Output:
	// Null: "null"
	// Ident: "TeSt"
	// String: "\"TeSt\""
	// Positive number: "12"
	// Negative number: "-12.23"
	// Boolean: "false"
	// List of capsule: "[value1,value2]"
	// Set of capsule: "[value1,value2]"
	// Object with capsule: "{\n\"A\"=A_value\n\"B\"=\"B_value\"\n}"
	// Map of capsule: "{\n\"A\"=A_value\n\"B\"=B_value\n}"
	// Tuple with capsule: "[\"A_value\",B_value,2]"
}

func ExampleSplitIterable() {
	list, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.List(cty.String),
	)
	if err != nil {
		panic(err)
	}

	setOfCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("value1"),
			cty.StringVal("value2"),
		},
		cty.Set(cty.String),
	)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	mapOfCapsule, err := gocty.ToCtyValue(
		map[string]cty.Value{
			"A": cty.StringVal("A_value"),
			"B": cty.StringVal("B_value"),
		},
		cty.Map(cty.String),
	)
	if err != nil {
		panic(err)
	}

	tupleWithCapsule, err := gocty.ToCtyValue(
		[]cty.Value{
			cty.StringVal("A_value"),
			cty.NumberIntVal(2),
		},
		cty.Tuple([]cty.Type{cty.String, cty.Number}),
	)
	if err != nil {
		panic(err)
	}

	start, elements, end := tokens.SplitIterable(list)
	fmt.Printf(
		"List:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(setOfCapsule)
	fmt.Printf(
		"Set:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(objectWithCapsule)
	fmt.Printf(
		"Object:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(mapOfCapsule)
	fmt.Printf(
		"Map:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	start, elements, end = tokens.SplitIterable(tupleWithCapsule)
	fmt.Printf(
		"Tuple:\n\tStart: %#v\n\tElements: %#v\n\tEnd: %#v\n",
		string(start.Bytes()),
		string(elements.Bytes()),
		string(end.Bytes()),
	)

	// Output:
	// List:
	// 	Start: "["
	// 	Elements: "\"value1\", \"value2\""
	// 	End: "]"
	// Set:
	// 	Start: "["
	// 	Elements: "\"value1\", \"value2\""
	// 	End: "]"
	// Object:
	// 	Start: "{"
	// 	Elements: "\n  A = 2\n  B = \"B_value\"\n"
	// 	End: "}"
	// Map:
	// 	Start: "{"
	// 	Elements: "\n  A = \"A_value\"\n  B = \"B_value\"\n"
	// 	End: "}"
	// Tuple:
	// 	Start: "["
	// 	Elements: "\"A_value\", 2"
	// 	End: "]"
}
