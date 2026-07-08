package main

import (
	//"fmt"
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	inputstr := strings.Split(input, "\\n")

	var result strings.Builder
	for _, n := range inputstr {
		for row := 0; row < 8; row++ {
			for _, r := range n {
				result.WriteRune(banner[r][row])

			}
		}
	}
	return  inputstr
}
