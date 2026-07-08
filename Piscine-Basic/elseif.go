package main

import "fmt"

func esleIf() {
	time := 11      //variable declaration
	if time >= 16 { //if condition
		fmt.Println("Good Evening")
	} else if time < 16 && time >= 12 { //else if condition
		fmt.Println("Good Day")
	} else { //else condition
		fmt.Println("Good Morning")
	}

	var a = 15
	var b = 15
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a and b are equal")
	}
}
