package main

import (
	"fmt"
	"strings"
)

func modifyWord() {
	for {
		// prompt user to imput desired word
		var input string
		fmt.Print("Input a word: ")
		fmt.Scanln(&input)

		// check for empty input
		if input == "" {
			fmt.Println("Word can not be empty! try again")
			continue
		}
		// prompt user to enter desired operation
		var operation string
		fmt.Print("Enter operatio (lastChar / capitalize / deleteIndex): ")
		fmt.Scanln(&operation)

		// switch operation according to user input
		switch strings.ToLower(operation) {

		case "lastchar":
			fmt.Println(lastChar(input))

		case "capitalize":
			fmt.Println(capitalize(input))

		case "deleteindex":

			var index int
			fmt.Print("Enter index to be deleted: ")
			fmt.Scanln(&index)

			result, err := deleteIndex(input, index)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(result)

		default:
			fmt.Println("Invalid operation! try again")
		}

		var continuation string
		fmt.Print("Would you like to modify another word? (yes/no): ")
		fmt.Scanln(&continuation)

		if strings.ToLower(continuation) != "yes" {
			fmt.Println("Bye! See ya later.")
			break
		}
	}
}

// Helper Functions
func lastChar(s string) string {
	runes := []rune(s)
	return string(runes[len(runes)-1])
}

func capitalize(s string) string {
	return strings.ToUpper(s)
}

func deleteIndex(s string, index int) (string, error) {

	runes := []rune(s)

	// for edge cases where inputed index is
	if index < 0 || index >= len(runes) {
		return "", fmt.Errorf("Invalid index %d for word of length %d", index, len(runes))
	}
	return string(runes[:index]) + string(runes[index+1:]), nil

}
