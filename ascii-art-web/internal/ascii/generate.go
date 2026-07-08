package ascii

import (
	"strings"
)
func BuildArt(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}

	err := ValidInput(text)
	if err != nil {
		return "invalid ascii character"
	}

	word := SplitText(text)

	var result strings.Builder
	for i, char := range word {
		if char == "" {
			if i == len(word)-1 {
				continue
			}
			result.WriteByte('\n')
		}
		output := RenderLine(char, banner)
		for _, value := range output {
			result.WriteString(value)
			result.WriteString("\n")
		}

	}

	return result.String()
}
