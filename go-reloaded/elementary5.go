package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) string {
	
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file")
		(os.Exit(1))
	}
	return string(file)

}
