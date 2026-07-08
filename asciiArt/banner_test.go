// package main

// import (
// 	"os"

// 	"strings"

// 	"testing"
// )

// // TestLoadBanner_ReturnsCorrectCharCount checks that all 95 printable ASCII

// // characters (codes 32–126) are present in the loaded banner map.

// func TestLoadBanner_ReturnsCorrectCharCount(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	if len(banner) != 95 {

// 		t.Errorf("expected 95 characters in banner, got %d", len(banner))

// 	}

// }

// // TestLoadBanner_EachCharHasEightLines checks that every character in the

// // banner map has exactly 8 art lines.

// func TestLoadBanner_EachCharHasEightLines(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	for r, art := range banner {

// 		if len(art) != 8 {

// 			t.Errorf("character %q: expected 8 lines, got %d", r, len(art))

// 		}

// 	}

// }

// // TestLoadBanner_AllPrintableASCIIPresent checks every rune from 32 to 126

// // individually to make sure none are missing.

// func TestLoadBanner_AllPrintableASCIIPresent(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	for code := rune(32); code <= 126; code++ {

// 		if _, ok := banner[code]; !ok {

// 			t.Errorf("character %q (ASCII %d) is missing from banner", code, code)

// 		}

// 	}

// }

// // TestLoadBanner_SpaceCharPresent specifically checks the space character

// // since it is invisible and students often accidentally skip it.

// func TestLoadBanner_SpaceCharPresent(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	art, ok := banner[' ']

// 	if !ok {

// 		t.Fatal("space character (ASCII 32) is missing from banner")

// 	}

// 	if len(art) != 8 {

// 		t.Errorf("space character: expected 8 lines, got %d", len(art))

// 	}

// }

// // TestLoadBanner_NoLineContainsNewline checks that none of the stored art

// // lines themselves contain a newline byte — the parser must strip them.

// func TestLoadBanner_NoLineContainsNewline(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	for r, art := range banner {

// 		for i, line := range art {

// 			if strings.Contains(line, "\n") {

// 				t.Errorf("character %q line %d contains a newline byte", r, i)

// 			}

// 		}

// 	}

// }

// // TestLoadBanner_KnownCharacters spot-checks a handful of characters that

// // have well-known first lines in the standard banner.
// /*
// func TestLoadBanner_KnownCharacters(t *testing.T) {

// 	banner, err := LoadBanner("standard.txt")

// 	if err != nil {

// 		t.Fatalf("unexpected error loading standard.txt: %v", err)

// 	}

// 	cases := []struct {
// 		char rune

// 		lineIdx int

// 		contains string
// 	}{

// 		{'A', 0, "_"}, // top of 'A' has an underscore

// 		{'a', 0, ""}, // first line of lowercase a is blank in standard

// 		{'Z', 0, "_"},

// 		{'!', 0, "_"},
// 	}

// 	for _, tc := range cases {

// 		art, ok := banner[tc.char]

// 		if !ok {

// 			t.Errorf("character %q not found in banner", tc.char)

// 			continue

// 		}

// 		if len(art) <= tc.lineIdx {

// 			t.Errorf("character %q: not enough lines (need index %d)", tc.char, tc.lineIdx)

// 			continue

// 		}

// 		if tc.contains != "" && !strings.Contains(art[tc.lineIdx], tc.contains) {

// 			t.Errorf("character %q line %d: expected to contain %q, got %q",

// 				tc.char, tc.lineIdx, tc.contains, art[tc.lineIdx])

// 		}

// 	}

// }
// */
// // TestLoadBanner_FileNotFound expects a non-nil error and a nil map when the

// // file does not exist.

// func TestLoadBanner_FileNotFound(t *testing.T) {

// 	banner, err := LoadBanner("notfound.txt")

// 	if err == nil {

// 		t.Error("expected a non-nil error for missing file, got nil")

// 	}

// 	if banner != nil {

// 		t.Error("expected nil map on error, got non-nil")

// 	}

// }

// // TestLoadBanner_EmptyFile expects an error when the file exists but is empty.

// func TestLoadBanner_EmptyFile(t *testing.T) {

// 	f, err := os.CreateTemp("", "empty_banner_*.txt")

// 	if err != nil {

// 		t.Fatal(err)

// 	}

// 	defer os.Remove(f.Name())

// 	f.Close()

// 	_, err = LoadBanner(f.Name())

// 	if err == nil {

// 		t.Error("expected error for empty file, got nil")

// 	}

// }

// // TestLoadBanner_InvalidContent expects an error when the file has random

// // content that does not conform to the 8-lines-per-character format.

// func TestLoadBanner_InvalidContent(t *testing.T) {

// 	f, err := os.CreateTemp("", "invalid_banner_*.txt")

// 	if err != nil {

// 		t.Fatal(err)

// 	}

// 	defer os.Remove(f.Name())

// 	_, _ = f.WriteString("this is not a valid banner file\nonly two lines\n")

// 	f.Close()

// 	_, err = LoadBanner(f.Name())

// 	if err == nil {

// 		t.Error("expected error for invalid banner content, got nil")

// 	}

// }

// // TestLoadBanner_ShadowAndThinkertoy verifies the other two banner files also

// // load correctly with 95 chars each having 8 lines.

// func TestLoadBanner_ShadowAndThinkertoy(t *testing.T) {

// 	for _, filename := range []string{"shadow.txt", "thinkertoy.txt"} {

// 		banner, err := LoadBanner(filename)

// 		if err != nil {

// 			t.Errorf("%s: unexpected error: %v", filename, err)

// 			continue

// 		}

// 		if len(banner) != 95 {

// 			t.Errorf("%s: expected 95 chars, got %d", filename, len(banner))

// 		}

// 		for r, art := range banner {

// 			if len(art) != 8 {

// 				t.Errorf("%s: character %q has %d lines, expected 8", filename, r, len(art))

// 			}

// 		}

// 	}

// }
package main

import (
	"os"
	"testing"
)

// helper: write a temp banner file and return its path
func writeTempBanner(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "banner_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name()
}

// helper: build 8-line character block
func makeChar(ch string) string {
	block := ""
	for i := 0; i < 8; i++ {
		block += ch + "\n"
	}
	return block
}

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := LoadBanner("nonexistent_file.txt")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	path := writeTempBanner(t, "")
	_, err := LoadBanner(path)
	if err == nil {
		t.Fatal("expected error for empty file, got nil")
	}
}

func TestLoadBanner_OnlyBlankLines(t *testing.T) {
	path := writeTempBanner(t, "\n\n\n\n")
	_, err := LoadBanner(path)
	if err == nil {
		t.Fatal("expected error for blank-only file, got nil")
	}
}

func TestLoadBanner_SingleCharacter(t *testing.T) {
	// space (ASCII 32) is the first expected character
	content := makeChar("line")
	path := writeTempBanner(t, content)
	banner, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := banner[' ']; !ok {
		t.Error("expected space character (32) in banner map")
	}
}

func TestLoadBanner_MultipleCharacters(t *testing.T) {
	// space = 32, '!' = 33
	content := makeChar("space") + "\n" + makeChar("exclaim")
	path := writeTempBanner(t, content)
	banner, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, ch := range []rune{' ', '!'} {
		if _, ok := banner[ch]; !ok {
			t.Errorf("expected rune %q in banner map", ch)
		}
	}
}

func TestLoadBanner_CorrectLineCount(t *testing.T) {
	content := makeChar("row")
	path := writeTempBanner(t, content)
	banner, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines, ok := banner[' ']
	if !ok {
		t.Fatal("space character not found in banner")
	}
	if len(lines) != 8 {
		t.Errorf("expected 8 lines per character, got %d", len(lines))
	}
}

func TestLoadBanner_IncorrectLineCount(t *testing.T) {
	// only 5 lines instead of 8
	content := "a\nb\nc\nd\ne\n"
	path := writeTempBanner(t, content)
	_, err := LoadBanner(path)
	if err == nil {
		t.Fatal("expected error for wrong line count, got nil")
	}
}

func TestLoadBanner_LineContent(t *testing.T) {
	lines := [8]string{"_", "| |", "| |", "|_|", "   ", "   ", "   ", "   "}
	content := ""
	for _, l := range lines {
		content += l + "\n"
	}
	path := writeTempBanner(t, content)
	banner, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, ok := banner[' ']
	if !ok {
		t.Fatal("space character not found in banner")
	}
	for i, expected := range lines {
		if got[i] != expected {
			t.Errorf("line %d: expected %q, got %q", i, expected, got[i])
		}
	}
}

func TestLoadBanner_AllPrintableASCII(t *testing.T) {
	content := ""
	for ch := 32; ch <= 126; ch++ {
		for i := 0; i < 8; i++ {
			content += "x\n"
		}
		if ch < 126 {
			content += "\n"
		}
	}
	path := writeTempBanner(t, content)
	banner, err := LoadBanner(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for ch := 32; ch <= 126; ch++ {
		if _, ok := banner[rune(ch)]; !ok {
			t.Errorf("missing rune %q (ASCII %d)", rune(ch), ch)
		}
	}
	if len(banner) != 95 {
		t.Errorf("expected 95 entries (32–126), got %d", len(banner))
	}
}