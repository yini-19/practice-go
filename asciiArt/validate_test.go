package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantRune rune
	}{
		{
			name:     "valid ascii string",
			input:    "Hello, World!",
			wantErr:  false,
			wantRune: 0,
		},
		{
			name:     "empty string",
			input:    "",
			wantErr:  false,
			wantRune: 0,
		},
		{
			name:     "contains newline character",
			input:    "Hello\nWorld",
			wantErr:  true,
			wantRune: '\n',
		},
		{
			name:     "contains unicode character",
			input:    "Hello😊",
			wantErr:  true,
			wantRune: '😊',
		},
		{
			name:     "contains tab character",
			input:    "Hello\tWorld",
			wantErr:  true,
			wantRune: '\t',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRune, err := Validate(tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("did not expect error, got %v", err)
			}

			if gotRune != tt.wantRune {
				t.Errorf("expected rune %q, got %q", tt.wantRune, gotRune)
			}
		})
	}
}
