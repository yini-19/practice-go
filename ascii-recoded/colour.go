package main

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
