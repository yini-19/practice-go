package main

import "fmt"

func timetable() {
	for i := 1; i <= 12; i++ {
		for j := 1; j <= 12; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
