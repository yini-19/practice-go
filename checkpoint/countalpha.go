package piscine

func CountAlpha(s string) int {
	count := 0
	for _, s := range s {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
			count++
		}
	}
	return count
}
