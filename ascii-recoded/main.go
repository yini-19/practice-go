package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [substring] [STRING]")
		return
	}
	inputString := os.Args[1]
	filename := "standard.txt"
	if len(os.Args) == 3 {
		//inputString = os.Args[1]
		filename = os.Args[2]
	} else {
		option := os.Args[1]
		//substring := os.Args[2]
		inputString = os.Args[3]
		color := strings.TrimPrefix(option, "--color=")

		if !strings.HasPrefix(option, "--color=") {
			fmt.Println("Usage: go run . [OPTION] [substring] [STRING]")
			return
		}
		color = GetColor(color)
		if len(os.Args) == 5 {
			filename = os.Args[4]
		}
	}

	banner, err := LoadBanner(filename)
	if err != nil {
		fmt.Printf("Error loading banner from %s: %v\n", filename, err)
		return
	}
	fmt.Print(GenerateArt(inputString, banner))

}
