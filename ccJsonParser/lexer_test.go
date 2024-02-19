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
			name: "Empty Input",
			inputFilePath: "tests/step1/invalid.json",
			expectedValues: []Token{
				{Type: EOF, Value: ""},
			},
		},
		{
			name: "Empty Object",
			inputFilePath: "tests/step1/valid.json",
			expectedValues: []Token{
				{Type: BeginObject, Value: "{"},
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
					t.Errorf("Expected %s but got %s", expected, actual)
				}
			}
		})
	}
}

