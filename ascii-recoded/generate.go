package main

import (
	//"image/color"
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

	words := SplitInput(input)

	for i, word := range words {
		if word == "" {
			if i == len(words)-1 && strings.HasSuffix(input, "\\n") {
				output.WriteString("")
				continue
			}
			output.WriteString("\n")
			continue
		}

		reset := "\033[0m"
		option := os.Args[1]
		color := strings.TrimPrefix(option, "--color=")
		color = GetColor(color)
		
		generate := RenderLine(word, banner, color, reset)
		for _, rendered := range generate {
			output.WriteString(rendered)
			output.WriteString("\n")
		}

	}
	return output.String()
}

// func GenerateArt(input string, banner map[rune][]string) string {
// 	if input == "" {
// 		return ""
// 	}
// 	if input == "\\n" {
// 		return "\n"
// 	}

// 	replace := strings.ReplaceAll(input, "\\n", "")
// 	words := SplitInput(input)
// 	var output []string

// 	for i, word := range words {
// 		if word == "" {
// 			if i == len(words)-1 && strings.HasSuffix(input, "\\n") {
// 				for i := 0; i < 8; i++ {
// 					output = append(output, "")
// 				}
// 				continue
// 			}
// 			output = append(output, "")
// 			continue
// 		}
// 		output = append(output, RenderLine(word, banner)...)
// 	}
// 	joined := strings.Join(output, "\n")
// 	if replace != "" {
// 		joined += "\n"
// 	}
// 	return joined
// }
