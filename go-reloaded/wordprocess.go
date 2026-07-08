package main

import (
	"fmt"
)

func process() {
	var input string
	var operation string
	var continuation string
	fmt.Print("Input a word: ")
	fmt.Scanln(&input)
	fmt.Print("Enter numbers 1-3 for operation to be carried out\n 1. lastChar\n 2. capitalize\n 3. deleteIndex\n Operation: ")
	fmt.Scanln(&operation)
	if operation == "1" {
		fmt.Println(lastChar(input))
	}
	if operation == "2" {
		fmt.Println(capitalize(input))
	}
	if operation == "3" {
		var index int
		fmt.Print("Select index to be deleted: ")
		fmt.Scanln(&index)
		if index >= len(input) {
			fmt.Println("Invalid Index! Start again")
			process()
		}
		fmt.Println(deleteIndex(input, index))

	}
	fmt.Print("Would like to modify another word: ")
	fmt.Scanln(&continuation)
	if continuation == "yes" {
		process()
	} else {
		fmt.Println("Bye! See ya later")
	}

}
