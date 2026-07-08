package cd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) (string, error) {
	if input == "" {
		return "", nil
	}
	if input == "\\n" {
		return "\n", nil
	}
	var output strings.Builder

	// validate input
	_, err := Validate(input)
	if err != nil {
		return "", err
	}

	// if input is valid, split by newline
	text := strings.ReplaceAll(input, "\r\n", "\n")
	words := strings.Split(text, "\n")
	for i, word := range words {
		if word == "" {
			if i == len(words)-1 && strings.HasSuffix(input, "\n") {
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
	return output.String(), err
}

func Validate(input string) (rune, error) {
	for _, char := range input {
		if char == '\n' || char == '\r' {
			continue
		}
		if char > 126 || char < 32 {
			return char, fmt.Errorf("Error! %v is not valid", char)
		}
	}
	return 0, nil

}

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	banner := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var lines []string
	charcode := 32

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if len(lines) > 0 {
				// Each character must have exactly 8 lines
				if len(lines) != 8 {
					return nil, os.ErrInvalid
				}
				banner[rune(charcode)] = lines
				lines = []string{}
				charcode++
				// Ensure charcode stays in valid ASCII printable range
				if charcode > 126 {
					return nil, os.ErrInvalid
				}
			}
			continue
		}
		lines = append(lines, line)
	}

	// check scanning error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// handle last character block if file didn't end with an empty separator
	if len(lines) > 0 {
		if len(lines) != 8 {
			return nil, os.ErrInvalid
		}
		banner[rune(charcode)] = lines
	}
	return banner, nil
}

func RenderLine(input string, banner map[rune][]string) []string {
	var builder [8]strings.Builder

	for _, char := range input {

		line, ok := banner[char]
		if !ok {
			log.Printf("Warning: character %q not found in banner, skipping\n", char)
			continue
		}
		if len(line) != 8 {
			log.Printf("Warning: character %q has malformed banner (expected 8 lines, got %d), skipping\n", char, len(line))
			continue
		}
		for i := 0; i < 8; i++ { // Assuming each character's banner is 8 lines tall
			builder[i].WriteString(line[i])
		}
	}

	var result []string
	for _, b := range builder {
		result = append(result, b.String())
	}
	return result
}
