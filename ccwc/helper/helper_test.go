package helper

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

var filepath string = "../test.txt"

func init() {
	_, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	} 
}

func TestHelper(t *testing.T) {	
	b := ReadFromFileOrStdin(filepath)

	type testCase struct {
		name string
		flagOption bool
		args []string
		actualOutputFn func(b []byte) int
	}

	testCases := []testCase{
		{
			name: "get byte count",
			flagOption: true,
			args: []string{"wc", "-c", filepath},
			actualOutputFn: GetByteCount,
		},
		{
			name: "get line count",
			flagOption: true,
			args: []string{"wc", "-l", filepath},
			actualOutputFn: GetLineCount,
		},
		{
			name: "get word count",
			flagOption: true,
			args: []string{"wc", "-w", filepath},
			actualOutputFn: GetWordCount,
		},
		{
			name: "get character count",
			flagOption: true,
			args: []string{"wc", "-m", filepath},
			actualOutputFn: GetCharacterCount,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.flagOption {
				actualOutput := getActualOutput(tc.actualOutputFn, b)
				expectedOutput := getExpectedOutput(tc.args)
				if actualOutput != expectedOutput {
					t.Errorf("Expected %d but got %d", expectedOutput, actualOutput)
				}
			}
		})
	}
}

func getExpectedOutput(args []string) int {
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Split(strings.Trim(string(output), " "), " ")[0]
	var num int
	num, err = strconv.Atoi(result)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func getActualOutput(fn func(b []byte) int, b []byte) int{
	return fn(b)
}