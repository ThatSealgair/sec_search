package censys

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ThatSealgair/sec_search/internal/platform/mappings"
	"github.com/ThatSealgair/sec_search/internal/platform/utils"
)

type QueryConverter struct{}

func (c *QueryConverter) ToCensys(query string) string {
	return query
}

func (c *QueryConverter) ToShodan(query string) string {
	parts := strings.Split(query, " AND ")
	var shodanParts []string

	for _, part := range parts {
		if !strings.Contains(part, ":") {
			continue
		}

		key, value := utils.SplitKeyValue(part)
		if shodanKey, ok := mappings.CensysToShodan[key]; ok {
			needsQuotes := !strings.Contains(shodanKey, ".status") && !strings.Contains(shodanKey, "port")
			formattedValue := utils.FormatValue(value, needsQuotes)
			shodanParts = append(shodanParts, fmt.Sprintf("%s:%s", shodanKey, formattedValue))
		}
	}

	return strings.Join(shodanParts, " ")
}

func (c *QueryConverter) ToHunter(query string) string {
	parts := strings.Split(query, " AND ")
	values := url.Values{}

	for _, part := range parts {
		if !strings.Contains(part, ":") {
			continue
		}

		key, value := utils.SplitKeyValue(part)
		switch key {
		case "autonomous_system.organization":
			values.Set("company", value)
		case "services.http.response.html":
			if strings.Contains(value, "@") {
				values.Set("email", value)
			}
		}
	}

	return values.Encode()
}
