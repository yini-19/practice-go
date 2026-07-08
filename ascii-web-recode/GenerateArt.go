package main

import (
	"ascii-web-recode/cd"
	"fmt"
)

func generateAscii(bannerFile string, text string) (string, error) {
	banner, err := cd.LoadBanner("banners/" + bannerFile + ".txt")
	if err != nil {
		return "", fmt.Errorf("load banner %q: %w", bannerFile, err)
	}
	result, err := cd.GenerateArt(text, banner)
	if err != nil {
		return "", fmt.Errorf("generate art: %w", err)
	}
	return result, nil
}
