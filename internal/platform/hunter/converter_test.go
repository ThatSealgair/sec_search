package hunter

import (
	"strings"
	"testing"
)

func TestQueryConverter_ToShodan(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "basic company query",
			input: "company=Example+Corp",
			expected: []string{
				`org:"Example Corp"`,
			},
		},
		{
			name:  "domain query",
			input: "domain=example.com",
			expected: []string{
				`hostname:"example.com"`,
			},
		},
	}

	converter := &QueryConverter{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.ToShodan(tt.input)
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected result to contain %q, got %q", expected, result)
				}
			}
		})
	}
}
