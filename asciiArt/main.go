package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <inputstring> `OR` go run . <inputstring> <filename.txt>")
		return
	}

	inputString := os.Args[1]
	filename := "standard.txt"
	if len(os.Args) == 3 {
		filename = os.Args[2]
	}

	
	banner, err := LoadBanner(filename)
	if err != nil {
		fmt.Printf("Error loading banner from %s: %v\n", filename, err)
		return
	}
	
	fmt.Print(GenerateArt(inputString, banner))

}
