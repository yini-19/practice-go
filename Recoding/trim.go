package main

import "strings"

func TrimArtRows(rows []string) []string {
	trimmed := make([]string, len(rows))
	for i, row := range rows {
		
		trimmed[i] = strings.TrimRight(row, " ")
	} 
	return trimmed
}