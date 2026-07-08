package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}
	const reset = "\033[0m"
	option := os.Args[1]
	color := strings.TrimPrefix(option, "--color=")
	if len(os.Args) == 3 {
		text := os.Args[2]
		color := GetColor(color)
		colored := color + text + reset
		fmt.Println(colored)
		return
	}
	substr := os.Args[2]
	text := os.Args[3]

	if !strings.HasPrefix(option, "--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	if strings.HasPrefix(option, "--color=") {

		color := strings.TrimPrefix(option, "--color=")

		color = GetColor(color)
		colored := strings.ReplaceAll(text, substr, color+substr+reset)
		fmt.Println(colored)

	}

}
func GetColor(color string) string {
	switch color {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "blue":
		return "\033[34m"
	default:
		return "\033[0m"
	}
}
