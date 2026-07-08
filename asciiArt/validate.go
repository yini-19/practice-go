package main

import "fmt"

func Validate(input string) (rune, error) {
	for _, char := range input {
		if char > rune(126) || char < rune(32) {
			return char, fmt.Errorf("Error! %v is not valid", string(char))
		}
	}
	return 0, nil

}
