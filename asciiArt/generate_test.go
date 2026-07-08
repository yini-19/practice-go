package main

import (
	"testing"
)

// mock banner for testing (simple 2-line height instead of 8 for simplicity)
func mockBanner() map[rune][]string {
	return map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': {"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
		' ': {"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
	}
}

// mock renderLines to isolate Generate behavior

func TestGenerate_SingleLine(t *testing.T) {
	banner := mockBanner()
	input := "AB"

	got := GenerateArt(input, banner)

	expected := "" +
		"A1B1\n" +
		"A2B2\n" +
		"A3B3\n" +
		"A4B4\n" +
		"A5B5\n" +
		"A6B6\n" +
		"A7B7\n" +
		"A8B8\n"

	if got != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, got)
	}
}

func TestGenerate_WithSpace(t *testing.T) {
	banner := mockBanner()
	input := "A B"

	got := GenerateArt(input, banner)

	expected := "" +
		"A1  B1\n" +
		"A2  B2\n" +
		"A3  B3\n" +
		"A4  B4\n" +
		"A5  B5\n" +
		"A6  B6\n" +
		"A7  B7\n" +
		"A8  B8\n"

	if got != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, got)
	}
}

func TestGenerate_MultipleLines(t *testing.T) {
	banner := mockBanner()
	input := "A\\nB"

	got := GenerateArt(input, banner)

	expected := "" +
		"A1\n" +
		"A2\n" +
		"A3\n" +
		"A4\n" +
		"A5\n" +
		"A6\n" +
		"A7\n" +
		"A8\n" +
		"B1\n" +
		"B2\n" +
		"B3\n" +
		"B4\n" +
		"B5\n" +
		"B6\n" +
		"B7\n" +
		"B8\n"

	if got != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, got)
	}
}
