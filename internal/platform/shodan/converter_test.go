package shodan

import (
	"strings"
	"testing"
)

func TestQueryConverter_ToCensys(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // Parts that should be in the output
	}{
		{
			name:  "basic network query",
			input: "port:443 country:US",
			expected: []string{
				"services.port: 443",
				`location.country_code: "US"`,
			},
		},
		{
			name:  "web service query",
			input: `product:"nginx" version:"1.18.0" http.status:200`,
			expected: []string{
				`services.software.product: "nginx"`,
				`services.software.version: "1.18.0"`,
				"services.http.response.status_code: 200",
			},
		},
		{
			name:  "ssl certificate query",
			input: `ssl:true ssl.cert.subject.cn:"*.example.com"`,
			expected: []string{
				`services.tls.certificates.leaf_data.subject.common_name: "*.example.com"`,
			},
		},
	}

	converter := &QueryConverter{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.ToCensys(tt.input)
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected result to contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestQueryConverter_ToHunter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "company query",
			input: `org:"Example Corp"`,
			expected: map[string]string{
				"company": "Example Corp",
			},
		},
		{
			name:  "domain query",
			input: `hostname:"example.com"`,
			expected: map[string]string{
				"domain": "example.com",
			},
		},
		{
			name:  "email query",
			input: `html:"user@example.com"`,
			expected: map[string]string{
				"email": "user@example.com",
			},
		},
	}

	converter := &QueryConverter{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.ToHunter(tt.input)
			for key, value := range tt.expected {
				if !strings.Contains(result, key+"="+value) {
					t.Errorf("Expected result to contain %s=%s, got %s", key, value, result)
				}
			}
		})
	}
}
