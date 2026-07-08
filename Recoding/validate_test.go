package main

import (
    //"strings"
    "testing"
)

// buildGoodBanner constructs a perfectly valid banner map in memory.
// No file needed — all 95 printable ASCII chars, each with exactly 8 lines.
func buildGoodBanner() map[rune][]string {
    banner := make(map[rune][]string)
    for r := rune(32); r <= 126; r++ {
        art := make([]string, 8)
        for i := range art {
            art[i] = string(r) + "row"
        }
        banner[r] = art
    }
    return banner
}

func TestValidateBanner_GoodMap(t *testing.T) {
    err := ValidateBanner(buildGoodBanner())
    if err != nil {
        t.Errorf("expected nil for valid banner, got: %v", err)
    }
}

func TestValidateBanner_Nil(t *testing.T) {
    err := ValidateBanner(nil)
    if err == nil {
        t.Error("expected error for nil banner, got nil")
    }
}

func TestValidateBanner_WrongEntryCount(t *testing.T) {
    banner := buildGoodBanner()
    delete(banner, 'A')
    delete(banner, 'B')
    err := ValidateBanner(banner)
    if err == nil {
        t.Error("expected error for banner with 93 entries, got nil")
    }
}
