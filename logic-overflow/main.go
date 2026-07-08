package main

import (
	"fmt"
	"os"
	"strings"
)

func Process(data []byte) string {
	return strings.ToUpper(string(data))
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// we will be using ReadFile and Args from the os package

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading Reading ")
		return
	}

	result := Process(data)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file")
	}
}
