package main

import (
	"reflect"
	"testing"
)

func TestSplitInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single line",
			input:    "Hello World",
			expected: []string{"Hello World"},
		},
		{
			name:     "multiple lines",
			input:    "Hello\\nWorld\\nGo",
			expected: []string{"Hello", "World", "Go"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "string starts with newline",
			input:    "\\nHello",
			expected: []string{"", "Hello"},
		},
		{
			name:     "string ends with newline",
			input:    "Hello\\n",
			expected: []string{"Hello", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitInput(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
