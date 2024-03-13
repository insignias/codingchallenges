package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "field number")
	delimiter := flag.String("d", "	", "delimiter")

	flag.Parse()

	intFields := formatFields(fields)
	// var fs []string
	// if strings.Contains(*fields, ",") {
	// 	fs = strings.Split(*fields, ",")
	// } else {
	// 	fs = strings.Split(*fields, " ")
	// }

	

	// for _, f := range fs {
	// 	num, err := strconv.Atoi(f)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	intFields = append(intFields, num)
	// }

	filename := flag.Arg(0)

	file := Readfile(filename)
	defer file.Close()

	scanner := GetNewScanner(file)

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

func Readfile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func GetNewScanner(file *os.File) *bufio.Scanner {
	return bufio.NewScanner(file)
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

	for _, line := range lines {
		columns := strings.Split(line, delimiter)
		if !strings.Contains(line, delimiter) {
			// if idx == len(lines)-1 {
			// 	output += fmt.Sprint(columns[0] + "%")
			// } else {
				output += fmt.Sprintln(columns[0])
			// }
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
		output += fmt.Sprintln()
	}

	return output, err

}
