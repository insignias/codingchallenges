package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	field := flag.Int("f", 1, "field number")
	delimiter := flag.String("d", "	", "delimiter")

	flag.Parse()

	filename := flag.Arg(0)

	file := Readfile(filename)
	defer file.Close()
	
	scanner := GetNewScanner(file)

	result, err := GetFields(scanner, *field, *delimiter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)

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

func GetFields(scanner *bufio.Scanner, field int, delimiter string) (string, error) {
	var err error
	var results []string
	var delimiterValidationError bool
	
	for scanner.Scan() {
		columns := strings.Split(string(scanner.Text()), delimiter)

		if field < 1 {
			err = fmt.Errorf("field value may not be zero")
		} 
		
		if len(columns) > 1 {
			if (field-1) < len(columns) {
				results = append(results, columns[field-1])
			} else {
				results = append(results, " ")
			}
		} else {
			results = append(results, columns...)
			delimiterValidationError = true
		}
	}

	
	var final string
	for idx, result := range results {
		if idx < len(results)-1 {
			final += result + "\n"
		} else {
			final += result 
			if !delimiterValidationError {
				final += "\n"
			}
		}
	}
	
	return final, err
}