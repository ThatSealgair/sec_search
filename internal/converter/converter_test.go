package converter

import (
	"testing"

	"github.com/ThatSealgair/sec_search/internal/platform"
)

func TestConverter_Convert(t *testing.T) {
	tests := []struct {
		name         string
		query        string
		fromPlatform platform.Platform
		toPlatform   platform.Platform
		wantErr      bool
		expectedPart string
	}{
		{
			name:         "shodan to censys - network",
			query:        "port:443 country:US",
			fromPlatform: platform.Shodan,
			toPlatform:   platform.Censys,
			wantErr:      false,
			expectedPart: "services.port: 443",
		},
		{
			name:         "censys to shodan - software",
			query:        `services.software.product: "nginx"`,
			fromPlatform: platform.Censys,
			toPlatform:   platform.Shodan,
			wantErr:      false,
			expectedPart: `product:"nginx"`,
		},
		{
			name:         "hunter to shodan - company",
			query:        "company=Example+Corp",
			fromPlatform: platform.Hunter,
			toPlatform:   platform.Shodan,
			wantErr:      false,
			expectedPart: `org:"Example Corp"`,
		},
		{
			name:         "same platform",
			query:        "port:443",
			fromPlatform: platform.Shodan,
			toPlatform:   platform.Shodan,
			wantErr:      false,
			expectedPart: "port:443",
		},
	}

	conv := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := conv.Convert(tt.query, tt.fromPlatform, tt.toPlatform)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(result, tt.expectedPart) {
				t.Errorf("Convert() = %v, expected to contain %v", result, tt.expectedPart)
			}
		})
	}
}

func FuzzConverter_Convert(f *testing.F) {
	seeds := []string{
		"port:443 country:US",
		`product:"nginx" version:"1.18.0"`,
		"company=Example+Corp",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	conv := New()
	f.Fuzz(func(t *testing.T, input string) {
		// Test all platform combinations
		platforms := []platform.Platform{platform.Shodan, platform.Censys, platform.Hunter}
		for _, from := range platforms {
			for _, to := range platforms {
				result, err := conv.Convert(input, from, to)
				if err != nil {
					// Errors are okay, but panics are not
					continue
				}
				if result == "" && input != "" {
					t.Errorf("Empty result for non-empty input %q", input)
				}
			}
		}
	})
}
