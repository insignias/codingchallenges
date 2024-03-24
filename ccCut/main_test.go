package main

import (
	"bufio"
	"strings"
	"log"
	"os"
	"os/exec"
	"testing"
	"github.com/insignias/codingchallenges/ccCut/helper"
)

var filepath string = "files/sample.tsv"
var filepath2 string = "files/fourchords.csv"

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
		expectedOutputFn func(*bufio.Scanner, []int, string) (string, error)
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
		{
			name: "Test multiple fields",
			actualOutputArgs: []string{"cut", "-f2,3", "-d,", filepath},
			expectedArgs: []string{"2,3", ",", filepath},
			expectedOutputFn: GetFields,
		},
		{
			name: "Test invalid field",
			actualOutputArgs: []string{"cut", "-f6", "-d,", filepath},
			expectedArgs: []string{"6", ",", filepath},
			expectedOutputFn: GetFields,
		},
		{
			name: "Test field with spaces",
			actualOutputArgs: []string{"cut", "-f1 2", "-d,", filepath},
			expectedArgs: []string{"1 2", ",", filepath},
			expectedOutputFn: GetFields,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := getActualOutput(tc.actualOutputArgs)
			field := tc.expectedArgs[0]
			fields := formatFields(&field)
			b := helper.ReadFromFileOrStdin(tc.expectedArgs[2])
			scanner := GetNewScanner(b)

			expectedOutput, err := getExpectedOutput(tc.expectedOutputFn, scanner, fields, tc.expectedArgs[1])
			if err != nil {
				t.Error(err)
			}
			
			if strings.Trim(actualOutput, "\n%") != strings.Trim(expectedOutput, "\n%") {
				t.Errorf("Expected \n%s but got \n%s", expectedOutput, actualOutput)
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

func getExpectedOutput(fn func(scanner *bufio.Scanner, fields []int, delimiter string) (string, error), scanner *bufio.Scanner, fields []int, delimiter string) (string, error) {
	return fn(scanner, fields, delimiter)
}