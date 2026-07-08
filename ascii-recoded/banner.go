package main

import (
	"bufio"
	//"fmt"
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
	if err := scanner.Err(); err != nil {
        return nil, err
    }
	if charcode > 126 && charcode < 32 || len(lines) != 8 {
		return nil, os.ErrInvalid
	}
	if len(lines) > 0 {
		banner[rune(charcode)] = lines
	}
	return banner, nil
}
