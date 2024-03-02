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

	flag.Parse()

	filename := flag.Arg(0)

	file := Readfile(filename)
	defer file.Close()
	
	scanner := GetNewScanner(file)

	result, err := GetFields(scanner, field)
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

func GetFields(scanner *bufio.Scanner, field *int) (string, error) {
	var err error
	var results []string
	for scanner.Scan() {
		columns := strings.Fields(string(scanner.Text()))
		if *field < 1 {
			err = fmt.Errorf("field value may not be zero")
		} 
		
		if (*field-1) < len(columns) {
			results = append(results, columns[*field-1])
		} else {
			results = append(results, " ")
		}
	}
	
	var final string
	for _, result := range results {
		final += result + "\n"
	}

	return final, err
}