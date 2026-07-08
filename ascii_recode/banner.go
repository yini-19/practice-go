package main

import (
	//"bufio"
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(file) == 0 {
		return nil, errors.New("empty file")
	}
	lines := strings.Split(string(file), "\n")
	lines = lines[1:]
	if len(lines) != 855 {
		return nil, errors.New("invalid content")
	}
	result := make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {
		// start := int(i-32)*9
		// end := start+8
		// result[i] = lines[start:end]
		for row := 0; row < 8; row++ {
			result[i] = append(result[i], lines[(int(i-32)*9)+row])
		}

	}
	return result, nil
	// file, err := os.Open(filename)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()

	// banner := make(map[rune][]string)
	// scanner := bufio.NewScanner(file)

	// var lines []string
	// charcode := 32

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if line == "" && len(lines) > 0 {
	// 		banner[rune(charcode)] = lines
	// 		lines = []string{}
	// 		charcode++
	// 	}
	// 	if line != "" {
	// 		lines = append(lines, line)
	// 	}
	// }
	// if len(lines) > 0 {
	// 	banner[rune(charcode)] = lines
	// }
	// return banner, err
}
