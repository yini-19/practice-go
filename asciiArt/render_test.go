package main

import (
	"bytes"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestRenderBanner_MissingChar(t *testing.T) {
	banner := map[rune][]string{
		'A': {
			" A ",
			"A A",
			"AAA",
			"A A",
			"A A",
			"A A",
			"A A",
			"A A",
		},
	}
	output := captureOutput(func() {
		RenderLine("B", banner)
	})
	if output == "" {
		t.Errorf("expected warning and blank lines for missing character")
	}
}


func TestRenderBanner_MalformedBanner(t *testing.T) {
	banner := map[rune][]string{
		'A': {
			" A ",
			"A A",
			"AAA",
			"A A",
			"A A",
			"A A",
			"A A",
			// Only 7 lines instead of 8
		},
	}
	output := captureOutput(func() {
		RenderLine("A", banner)
	})
	if output == "" {
		t.Errorf("expected blank lines for malformed banner character")
	}
}

func TestRenderBanner_AllCharsMissing(t *testing.T) {
    banner := map[rune][]string{
        'A': {
            " A ",
            "A A",
            "AAA",
            "A A",
            "A A",
            "A A",
            "A A",
            "A A",
        },
    }
    expected := "\n\n\n\n\n\n\n\n"
    output := RenderLine("BC", banner)
    joined := ""
    for _, line := range output {
        joined += line + "\n"
    }
    if joined != expected {
        t.Errorf("expected all blank lines for missing chars, got:\n%q", joined)
    }
}

func TestRenderBanner_InterleavedMalformedAndMissing(t *testing.T) {
    banner := map[rune][]string{
        'A': {
            " A ",
            "A A",
            "AAA",
            "A A",
            "A A",
            "A A",
            "A A",
            // Only 7 lines instead of 8
        },
        'B': {
            "BB ",
            "B B",
            "BB ",
            "B B",
            "BB ",
            "B B",
            "B B",
            "BB ",
        },
    }
    // 'A' is malformed, 'C' is missing, only 'B' should render
    expected := "BB \nB B\nBB \nB B\nBB \nB B\nB B\nBB \n"
    output := RenderLine("ABC", banner)
    joined := ""
    for _, line := range output {
        joined += line + "\n"
    }
    if joined != expected {
        t.Errorf("expected only B to render, got:\n%q", joined)
    }
}

func TestRenderBanner_MultiCharOutput(t *testing.T) {
    banner := map[rune][]string{
        'X': {
            "X X",
            " X ",
            " X ",
            " X ",
            " X ",
            " X ",
            " X ",
            "X X",
        },
        'Y': {
            "Y Y",
            "Y Y",
            " Y ",
            " Y ",
            " Y ",
            " Y ",
            " Y ",
            " Y ",
        },
    }
    expected := "X XY Y\n X Y Y\n X  Y \n X  Y \n X  Y \n X  Y \n X  Y \nX X Y \n"
    output := RenderLine("XY", banner)
    joined := ""
    for _, line := range output {
        joined += line + "\n"
    }
    if joined != expected {
        t.Errorf("expected XY banner, got:\n%q", joined)
    }
}
