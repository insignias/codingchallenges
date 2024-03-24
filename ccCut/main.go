package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/insignias/codingchallenges/ccCut/helper"
)

func main() {
	fields := flag.String("f", "", "field number")
	delimiter := flag.String("d", "	", "delimiter")

	flag.Parse()

	intFields := formatFields(fields)

	filename := flag.Arg(0)

	b := helper.ReadFromFileOrStdin(filename)

	scanner := GetNewScanner(b)

	output, err := GetFields(scanner, intFields, *delimiter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}

func formatFields(fields *string) []int {
	var intFields []int
	var fs []string
	if strings.Contains(*fields, ",") {
		fs = strings.Split(*fields, ",")
	} else {
		fs = strings.Split(*fields, " ")
	}

	for _, f := range fs {
		num, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		intFields = append(intFields, num)
	}
	return intFields
}

func GetNewScanner(b []byte) *bufio.Scanner {
	reader := bytes.NewReader(b)
	return bufio.NewScanner(reader)
}

func GetFields(scanner *bufio.Scanner, fields []int, delimiter string) (string, error){
	// flag := false
	var err error
	var line string
	var lines []string
	var output string

	for scanner.Scan() {
		line = string(scanner.Text())
		lines = append(lines, line)
	}

	for idx, line := range lines {
		columns := strings.Split(line, delimiter)
		if !strings.Contains(line, delimiter) {
			output += fmt.Sprintln(columns[0])
			continue
		}

		var newFields []int

		for _, field := range fields {
			if field < 1 {
				err = fmt.Errorf("field value may not be zero")
			} else if !(field > len(columns)) {
				newFields = append(newFields, field)
			}
		}

		for _, field := range newFields {
			if field != newFields[len(newFields)-1] {
				output += fmt.Sprintf("%s%s", columns[field-1], delimiter)
			} else {
				output += fmt.Sprint(columns[field-1])
			}
		}

		if idx != len(lines)-1 {
			output += fmt.Sprintln()
		}
		
	}

	return output, err

}
