package censys

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
			name:  "basic service query",
			input: `services.port: 443 AND location.country_code: "US"`,
			expected: []string{
				"port:443",
				`country:"US"`,
			},
		},
		{
			name:  "software query",
			input: `services.software.product: "nginx" AND services.software.version: "1.18.0"`,
			expected: []string{
				`product:"nginx"`,
				`version:"1.18.0"`,
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
