package main

import (
	"strings"
	"unicode"
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}
func toLower(s string) string {
	return strings.ToLower(s)
}
func capitalise(s string) string {
	if len(s)== 0{
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
