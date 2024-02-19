package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParser(t *testing.T){
	testCases := []struct{
		name string
		inputFilePath string
		expectedValue interface{}
		expectedError bool
	} {
		{
			name: "step1 invalid",
			inputFilePath: "tests/step1/invalid.json",
			expectedValue: nil,
			expectedError: true,
		},
		{
			name: "step1 valid",
			inputFilePath: "tests/step1/valid.json",
			expectedValue: make(map[string]interface{}),
			expectedError: false,
		},
		{
			name: "step2 invalid",
			inputFilePath: "tests/step2/invalid.json",
			expectedValue: map[string]interface{}{"key": "value"},
			expectedError: true,
		},
		{
			name: "step2 invalid2",
			inputFilePath: "tests/step2/invalid2.json",
			expectedValue: map[string]interface{}{"key": "value"},
			expectedError: true,
		},
		{
			name: "step2 valid",
			inputFilePath: "tests/step2/valid.json",
			expectedValue: map[string]interface{}{"key": "value"},
			expectedError: false,
		},
		{
			name: "step2 valid2",
			inputFilePath: "tests/step2/valid2.json",
			expectedValue: map[string]interface{}{"key": "value","key2": "value"},
			expectedError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.inputFilePath)
			if err != nil {
				t.Error(err)
			}
			p := NewParser(string(b))

			actualValue, actualError := p.Parse()

			if actualError != nil && !tt.expectedError {
				t.Errorf("unexpected error %s", actualError)
			}
			if !reflect.DeepEqual(actualValue, tt.expectedValue){
				t.Errorf("expected %v but got %v", tt.expectedValue, actualValue)
			}
		})
	}
}

