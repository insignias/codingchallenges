package main

import (
	"os"
	"testing"
)

func TestLexer(t *testing.T) {
	testCases := []struct{
		name string
		inputFilePath string
		expectedValues []Token
	} {
		// {
		// 	name: "step1 invalid",
		// 	inputFilePath: "tests/step1/invalid.json",
		// 	expectedValues: []Token{
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step1 valid",
		// 	inputFilePath: "tests/step1/valid.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step2 invalid",
		// 	inputFilePath: "tests/step2/invalid.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step2 valid",
		// 	inputFilePath: "tests/step2/valid.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step2 invalid2",
		// 	inputFilePath: "tests/step2/invalid2.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: Illegal, Value: "key2"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step2 valid2",
		// 	inputFilePath: "tests/step2/valid2.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key2"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step3 valid",
		// 	inputFilePath: "tests/step3/valid.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key1"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: True, Value: "true"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key2"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: False, Value: "false"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key3"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: Null, Value: "null"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key4"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key5"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: Number, Value: "101"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		// {
		// 	name: "step3 invalid",
		// 	inputFilePath: "tests/step3/invalid.json",
		// 	expectedValues: []Token{
		// 		{Type: BeginObject, Value: "{"},
		// 		{Type: String, Value: "key1"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: True, Value: "true"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key2"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: Illegal, Value: "False"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key3"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: Null, Value: "null"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key4"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: String, Value: "value"},
		// 		{Type: ObjectSeparator, Value: ","},
		// 		{Type: String, Value: "key5"},
		// 		{Type: ValueSeparator, Value: ":"},
		// 		{Type: Number, Value: "101"},
		// 		{Type: EndObject, Value: "}"},
		// 		{Type: EOF, Value: ""},
		// 	},
		// },
		{
			name: "step4 invalid",
			inputFilePath: "tests/step4/invalid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-n"},
				{Type: ValueSeparator, Value: ":"},
				{Type: Number, Value: "101"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-o"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "inner key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "inner value"},
				{Type: EndObject, Value: "}"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-l"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginArray, Value: "["},
				{Type: Illegal, Value: "list value"},
				{Type: EndArray, Value: "]"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step4 valid",
			inputFilePath: "tests/step4/valid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-n"},
				{Type: ValueSeparator, Value: ":"},
				{Type: Number, Value: "101"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-o"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginObject, Value: "{"},
				{Type: EndObject, Value: "}"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-l"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginArray, Value: "["},
				{Type: EndArray, Value: "]"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step4 valid2",
			inputFilePath: "tests/step4/valid2.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-n"},
				{Type: ValueSeparator, Value: ":"},
				{Type: Number, Value: "101"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-o"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "inner key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "inner value"},
				{Type: EndObject, Value: "}"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key-l"},
				{Type: ValueSeparator, Value: ":"},
				{Type: BeginArray, Value: "["},
				{Type: String, Value: "list value"},
				{Type: EndArray, Value: "]"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.inputFilePath)
			if err != nil {
				t.Error(err)
			}
			l := Newlexer(string(b))
			for _, expected := range tt.expectedValues {
				actual := l.NextToken()
				if actual != expected {
					t.Errorf("expected %s but got %s", expected, actual)
				}
			}
		})
	}
}

