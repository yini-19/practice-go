package main

import "fmt"

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char > rune(126) || char < rune(32) {
			return char, fmt.Errorf("Error! %v is not valid", string(char))
		}
	}
	return 0, nil

}

func Validate(banner map[rune][]string) error {
	if len(banner) != 95 {
		fmt.Println("Error: incomplete banner file")
	}
	for key, value := range banner {
		if key < 32 || key > 126 {
			fmt.Println("Error. invalid character")
		}

		if len(value) != 8 {
			fmt.Println("error!")
		}
	}
	return nil
}
