package main

import (
	"bufio"
	"os"
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
		}
		if line != "" {
			lines = append(lines, line)
		}
	}
	if len(banner) == 0 && len(lines) == 0 {
		return nil, os.ErrInvalid
	}
	if charcode > 126 || charcode < 32 || len(lines) != 8 {
		return nil, os.ErrInvalid
	}
	if len(lines) > 0 {
		banner[rune(charcode)] = lines
	}
	return banner, nil
}

// func RenderLine(input string, banner map[rune][]string) []string {
// 	var builder [8]strings.Builder

// 	for _, char := range input {
// 		line, ok := banner[char]
// 		if !ok {
// 			fmt.Println("Error")
// 			continue
// 		}
// 		if len(line) != 8 {
// 			fmt.Println("Error!")
// 			continue
// 		}

// 		for i := 0; i < 8; i++ {
// 			builder[i].WriteString(line[i])
// 		}

// 	}

// 	var result []string

// 	for _, b := range builder {
// 		result = append(result, b.String())
// 	}
// 	return result
// }

// func SplitInput(input string) []string {
// 	words := strings.Split(input, "\\n")
// 	return words
// }

// func Validate(input string) (rune, error) {
// 	for _, char := range input {
// 		if char > rune(126) || char < rune(32) {
// 			return char, fmt.Errorf("Error! %s is not a valid character", char)
// 		}
// 	}
// 	return 0, nil
// }

// func Generate(input string, banner map[rune][]string) string {
// 	words := SplitInput(input)
// 	var output strings.Builder

// 	for i, word := range words {
// 		if word == "" {
// 			if i < len(words)-1 {
// 				output.WriteString("\n")
// 			}
// 		}
// 		for _, char := range word {

// 			if _, ok := banner[char]; !ok {
// 				fmt.Fprintf(os.Stderr, "Non ASCII character %q", char)
// 				os.Exit(1)
// 			}
// 		}
// 		rendered := RenderLine(word, banner)
// 		for _, render := range rendered {
// 			output.WriteString(render + "\n")
// 		}
// 	}
// 	return output.String()
// }
