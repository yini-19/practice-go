package main

import "strings"

func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(input, "\n")

}
