package main

func MergeBanners(base, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string, len(base) + len(priority))

	for key, value := range base {
		result[key] = append([]string(nil), value...)
	}

	for key, value := range priority {
		result[key] = append([]string(nil), value...)
	}
	return result
}