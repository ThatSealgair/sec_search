package utils

import (
	"strings"
	"testing"
)

func TestSplitQuery(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic query",
			input:    "port:443 country:US",
			expected: []string{"port:443", "country:US"},
		},
		{
			name:     "query with quotes",
			input:    `product:"nginx" version:"1.18.0"`,
			expected: []string{`product:"nginx"`, `version:"1.18.0"`},
		},
		{
			name:     "empty query",
			input:    "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitQuery(tt.input, " ")
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d parts, got %d", len(tt.expected), len(result))
			}
			for i, part := range result {
				if part != tt.expected[i] {
					t.Errorf("Part %d: expected %q, got %q", i, tt.expected[i], part)
				}
			}
		})
	}
}

func FuzzSplitQuery(f *testing.F) {
	seeds := []string{
		"port:443 country:US",
		`product:"nginx" version:"1.18.0"`,
		"ip:192.168.1.1",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		result := SplitQuery(input, " ")
		for _, part := range result {
			if !strings.Contains(input, part) {
				t.Errorf("Result contains part %q not present in input %q", part, input)
			}
		}
	})
}
