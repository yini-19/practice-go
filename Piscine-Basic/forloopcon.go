package main

import "fmt"

func forLoop() {
	for i := 0; i < 9; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

}
