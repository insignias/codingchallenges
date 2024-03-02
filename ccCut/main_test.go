package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"testing"
)

var filepath string = "tests/sample.tsv"

func init() {
	_, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	} 
}

func TestMain(t *testing.T) {
	file := Readfile(filepath)
	scanner := GetNewScanner(file)

	testCases := []struct{
		name string
		actualOutputArgs []string
		expectedArgs int
		expectedOutputFn func(*bufio.Scanner, *int) (string, error)
	} {
		{
			name: "Test valid field",
			actualOutputArgs: []string{"cut", "-f", filepath},
			expectedArgs: 2,
			expectedOutputFn: GetFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := getActualOutput(tc.actualOutputArgs)
			expectedOutput, err := getExpectedOutput(tc.expectedOutputFn, scanner, &tc.expectedArgs)
			if err != nil {
				t.Error(err)
			}
			if actualOutput != expectedOutput {
				t.Errorf("Expected %s but got %s", expectedOutput, actualOutput)
			}
			
		})
	}
}

func getActualOutput(args []string) string {
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func getExpectedOutput(fn func(scanner *bufio.Scanner, field *int) (string, error), scanner *bufio.Scanner, field *int) (string, error) {
	return fn(scanner, field)
}