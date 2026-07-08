package main

import "strings"

type Tag struct {
	Kind  string // "up", "low", "cap", "hex", "bin"
	Count int    // 0 means apply to 1 word
}

func parseTag(token string) (Tag, bool) {
	token = strings.TrimSpace(token)
	token= strings.Trim(token, "()")
	tokens := strings.SplitN(token, ",", 2)
}

// parseTag("(up)")     → {Kind:"up",  Count:0}, true
// parseTag("(low,3)")  → {Kind:"low", Count:3}, true
// parseTag("(cap, 6)") → {Kind:"cap", Count:6}, true
// parseTag("(hex)")    → {Kind:"hex", Count:0}, true
// parseTag("hello")    → {}, false
