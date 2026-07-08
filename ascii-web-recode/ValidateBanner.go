package main

import (
	"fmt"
	"net/url"
)

func validateBanner(values url.Values, file string) (string, error) {

	bannerValues, exist := values[file]
	if !exist || len(bannerValues) == 0 {
		return "", fmt.Errorf("banner field is required")
	}
	bannerFile := bannerValues[0]
	if bannerFile == "" {
		return "", fmt.Errorf("banner value cannot be empty")
	}
	if !allowed[bannerFile] {
		return "", fmt.Errorf("invalid banner: must be standard, shadow, or thinkertoy")
	}
	return bannerFile, nil
}
