package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}
	var output strings.Builder

	// validate input
	if _, err := Validate(input); err != nil {
		fmt.Fprintf(os.Stderr, "Input validation failed: %v\n", err)
		os.Exit(1)
	}

	// if input is valid, split by newline
	words := SplitInput(input)

	for i, word := range words {
		if word == "" {
			if i == len(words)-1 && strings.HasSuffix(input, "\\n") {
				output.WriteString("")
				continue
			}
			output.WriteString("\n") // Add an extra newline for empty lines
			continue                 // Skip rendering for empty words
		}

		// render word to ascii art representation
		generate := RenderLine(word, banner)
		for _, rendered := range generate {
			output.WriteString(rendered)
			output.WriteString("\n")
		}

	}
	return output.String()
}
