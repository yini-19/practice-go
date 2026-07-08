// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func LoadBanner(filename string) (map[rune][]string, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading file")
// 	}

// 	if len(data) == 0 {
// 		return nil, fmt.Errorf("error! empty banner file")
// 	}

// 	content := strings.ReplaceAll(string(data), "\r\n", "\n")
// 	rawfile := strings.Split(content, "\n\n")

// 	bannerMap := make(map[rune][]string)
// 	charcode := ' '

// 	for _, raw := range rawfile {
// 		lines := strings.Split(raw, "\n")
// 		if len(lines) < 8 {
// 			return nil, fmt.Errorf("invalid character")
// 		}
// 		bannerMap[charcode] = lines[:8]
// 		charcode++
// 	}
// 	return bannerMap, nil
// }
