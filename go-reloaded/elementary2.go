package main

import "strings"

func splitWords(s string) []string {
	return strings.Fields(s)
}
func joinWords(words []string) string {
	return strings.Join(words, " ")
}
