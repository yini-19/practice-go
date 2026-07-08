package main

import "fmt"

func nestedIf() {
	num := 30
	if num > 15 {
		fmt.Println("number is greater than 15")
		if num > 20 {
			fmt.Println("number is also greater than 20")
		}
	} else {
		fmt.Println("number is less than 10")
	}
}
