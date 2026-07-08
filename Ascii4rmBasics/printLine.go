package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exercise:
// Write a Go program that reads a text file and prints its contents line by line.

// 1

func printLines() {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// 2

func printLine() {
	file, err := os.Open("standard.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func printLiner() {
	file, err := os.Open("standard.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}

}
