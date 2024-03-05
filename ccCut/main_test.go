package main

import (
	"bufio"

	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var filepath string = "tests/sample.tsv"
var filepath2 string = "tests/fourchords.csv"

func init() {
	_, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	} 
}

func TestMain(t *testing.T) {

	testCases := []struct{
		name string
		actualOutputArgs []string
		expectedArgs []string
		expectedOutputFn func(*bufio.Scanner, int, string) (string, error)
	} {
		{
			name: "Test default delimiter",
			actualOutputArgs: []string{"cut", "-f2", filepath},
			expectedArgs: []string{"2", "	", filepath},
			expectedOutputFn: GetFields,
		},
		{
			name: "Test custom delimiter",
			actualOutputArgs: []string{"cut", "-f2", "-d,", filepath2},
			expectedArgs: []string{"2", ",", filepath2},
			expectedOutputFn: GetFields,
		},
		{
			name: "Test wrong delimiter",
			actualOutputArgs: []string{"cut", "-f2", "-d,", filepath},
			expectedArgs: []string{"2", ",", filepath},
			expectedOutputFn: GetFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := getActualOutput(tc.actualOutputArgs)
			field, err := strconv.Atoi(tc.expectedArgs[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			file := Readfile(tc.expectedArgs[2])
			scanner := GetNewScanner(file)

			expectedOutput, err := getExpectedOutput(tc.expectedOutputFn, scanner, field, tc.expectedArgs[1])
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

func getExpectedOutput(fn func(scanner *bufio.Scanner, field int, delimiter string) (string, error), scanner *bufio.Scanner, field int, delimiter string) (string, error) {
	return fn(scanner, field, delimiter)
}