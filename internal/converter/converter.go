// internal/converter/converter.go
package converter

import (
	"fmt"

	"github.com/ThatSealgair/sec_search/internal/platform"
	"github.com/ThatSealgair/sec_search/internal/platform/censys"
	"github.com/ThatSealgair/sec_search/internal/platform/hunter"
	"github.com/ThatSealgair/sec_search/internal/platform/shodan"
)

type QueryConverter struct {
	platforms map[platform.Platform]platform.Converter
}

func New() *QueryConverter {
	return &QueryConverter{
		platforms: map[platform.Platform]platform.Converter{
			platform.Shodan: &shodan.QueryConverter{},
			platform.Censys: &censys.QueryConverter{},
			platform.Hunter: &hunter.QueryConverter{},
		},
	}
}

func (qc *QueryConverter) Convert(query string, from, to platform.Platform) (string, error) {
	if from == to {
		return query, nil
	}

	converter, exists := qc.platforms[from]
	if !exists {
		return "", fmt.Errorf("unsupported source platform: %s", from)
	}

	switch to {
	case platform.Shodan:
		return converter.ToShodan(query), nil
	case platform.Censys:
		return converter.ToCensys(query), nil
	case platform.Hunter:
		return converter.ToHunter(query), nil
	default:
		return "", fmt.Errorf("unsupported target platform: %s", to)
	}
}
