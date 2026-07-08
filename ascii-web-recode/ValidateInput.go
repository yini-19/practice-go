package main

import (
	"ascii-web-recode/cd"
	"fmt"
	"net/url"
	"strings"
)

func validateInput(values url.Values, name string) (string, error) {

	textValues, exists := values[name]
	if !exists || len(textValues) == 0 {
		return "", fmt.Errorf("text field is required")
	}
	if len(textValues[0]) > maxInputLength {
		return "", fmt.Errorf("maximum length exceeded. input not more than 500 characters")
	}
	text := strings.TrimSpace(textValues[0])
	if text == "" {
		return "", fmt.Errorf("text input cannot be empty or whitespace")
	}
	if _, err := cd.Validate(text); err != nil {
		return "", fmt.Errorf("character validation failed: only printable ascii characters allowed")

	}
	return text, nil
}
