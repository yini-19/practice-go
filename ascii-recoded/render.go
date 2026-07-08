package main

import (
	"fmt"
	"os"
	"strings"
)

func RenderLine(input string, banner map[rune][]string, color, reset string) []string {
	var builder [8]strings.Builder

	substring := os.Args[2]
	if substring != "" && !strings.Contains(input, substring) {
		fmt.Printf("Warning: input %q does not contain substring %q, skipping\n", input, substring)
		return []string{}
	}

	for _, char := range input {
		line, ok := banner[char]

		if !ok {
			fmt.Printf("Warning: character %q not found in banner, skipping\n", char)
			continue
		}
		if len(line) != 8 {
			fmt.Printf("Warning: character %q has malformed banner (expected 8 lines, got %d), skipping\n", char, len(line))
			continue
		}
		for i := 0; i < 8; i++ {
			if strings.Contains(substring, string(char)) {
				//for i := 0; i < 8; i++ { // Assuming each character's banner is 8 lines tall
				builder[i].WriteString(color)
				builder[i].WriteString(line[i])
				builder[i].WriteString(reset)
				//}
			} else { // Assuming each character's banner is 8 lines tall
				builder[i].WriteString(line[i])
			}
		}

	}

	var result []string
	for _, b := range builder {
		result = append(result, b.String())
	}

	return result
}
