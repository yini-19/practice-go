package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
	}
	filename := os.Args[1] + ".txt"
	if strings.HasSuffix(os.Args[1], ".txt") {
		filename = os.Args[1]
	}

	s := []string{"Go", "is", "fun"}
	fmt.Println(toUpper("hello"))
	fmt.Println(toLower("HELLO"))
	fmt.Println(capitalise("world"))
	fmt.Printf("%q\n", splitWords("Hello world"))
	fmt.Printf("%q\n", splitWords("  Go  is  fun "))
	fmt.Println(joinWords(s))
	fmt.Println(binToDecimal("abc"))
	fmt.Println(hexToDecimal("zz"))
	fmt.Println(readFile(filename))

}
