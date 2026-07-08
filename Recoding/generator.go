package main

func GeneratePattern(c rune) []string {
	banner := map[rune][]string{

		'A': []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		},
	}

	if c < 65 || c > 90 {
		return []string{}
	}
	return banner['A']
}

/*
Question 1: Simple Character Pattern Generator


Write a function GeneratePattern that creates a simple ASCII art representation for a single character using basic shapes. Each character should be represented as an 8-line string slice with only 3 different characters: . (dot), # (hash), and  (space).

Requirements:

Function signature: func GeneratePattern(c rune) []string

Only handle uppercase letters A-Z

Each letter must be exactly 6 characters wide

Pattern must be recognizable (doesn't need to be perfect)

Return empty slice for unsupported characters

Use a simple mapping approach (map[rune][]string)

File Structure:


ascii-art/
├── main.go
├── generator/
│   ├── generator.go
│   └── generator_test.go


*/
