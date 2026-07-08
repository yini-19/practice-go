package main

import(
	"fmt"
	"os"
)
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <inputfile> <outputfile>")
		os.Exit(1)
	}
	filename := os.Args[1]
	outputfile := os.Args[2]

	_, err := os.Stat(filename)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %v | Output: %v", filename, outputfile)

}
