package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		if line == "" && len(lines) > 0 {
			banner[rune(charcode)] = lines
			lines = []string{}
			charcode++

			if charcode > 126 {
				return nil, os.ErrInvalid
			}
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, os.ErrInvalid
	}
	if len(lines) > 0 {
		if len(lines) != 8 {
			fmt.Println("incomplete banner detected")
		}
		banner[rune(charcode)] = lines
	}
	return banner, nil
}

func GenerateArt(input string, banner map[rune][]string) (string, error) {

	if input == "" {
		return "", nil
	}
	if input == "\\n" {
		return "\n", nil
	}

	var output strings.Builder
	input = strings.ReplaceAll(input, "\r\n", "\n")
	words := strings.Split(input, "\n")

	for i, word := range words {

		if word == "" {
			if i == len(words)-1 && strings.HasSuffix(input, "\n") {
				output.WriteString("")
				continue
			}
			output.WriteString("\n")
			continue
		}
		rendered := RenderLine(word, banner)
		for _, render := range rendered {
			output.WriteString(render)
			output.WriteString("\n")
		}
	}
	return output.String(), nil
}

func RenderLine(input string, banner map[rune][]string) []string {
	var builder [8]strings.Builder

	for _, char := range input {
		line, ok := banner[char]
		if !ok {
			fmt.Println("warning, missing banner detected")
			continue
		}
		if len(line) != 8 {
			fmt.Println("warning, malformed banner detected")
			continue
		}

		for i := 0; i < 8; i++ {
			builder[i].WriteString(line[i])
		}
	}
	var result []string
	for _, b := range builder {
		result = append(result, b.String())
	}
	return result
}
