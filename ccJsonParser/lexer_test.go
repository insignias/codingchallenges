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
		{
			name: "step1 invalid",
			inputFilePath: "tests/step1/invalid.json",
			expectedValues: []Token{
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step1 valid",
			inputFilePath: "tests/step1/valid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step2 invalid",
			inputFilePath: "tests/step2/invalid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step2 valid",
			inputFilePath: "tests/step2/valid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step2 invalid2",
			inputFilePath: "tests/step2/invalid2.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: Illegal, Value: "k"},
				{Type: Illegal, Value: "e"},
				{Type: Illegal, Value: "y"},
				{Type: Illegal, Value: "2"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: EndObject, Value: "}"},
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "step2 valid2",
			inputFilePath: "tests/step2/valid2.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
				{Type: String, Value: "key"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
				{Type: ObjectSeparator, Value: ","},
				{Type: String, Value: "key2"},
				{Type: ValueSeparator, Value: ":"},
				{Type: String, Value: "value"},
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
					t.Errorf("expected %s bg %s", expected, actual)
				}
			}
		})
	}
}

