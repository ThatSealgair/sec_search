package shodan

import (
	"fmt"
	"strings"

	"github.com/ThatSealgair/sec_search/internal/platform/mappings"
	"github.com/ThatSealgair/sec_search/internal/platform/utils"
)

type QueryConverter struct{}

func (c *QueryConverter) ToShodan(query string) string {
	return query
}

func (c *QueryConverter) ToCensys(query string) string {
	parts := utils.SplitQuery(query, " ")
	var censysParts []string

	for _, part := range parts {
		if !strings.Contains(part, ":") {
			continue
		}

		key, value := utils.SplitKeyValue(part)
		if censysKey, ok := mappings.ShodanToCensys[key]; ok {
			needsQuotes := !strings.Contains(key, ".status") && !strings.Contains(key, "port")
			formattedValue := utils.FormatValue(value, needsQuotes)
			censysParts = append(censysParts, fmt.Sprintf("%s: %s", censysKey, formattedValue))
		}
	}

	return strings.Join(censysParts, " AND ")
}

func (c *QueryConverter) ToHunter(query string) string {
	parts := utils.SplitQuery(query, " ")
	var queryParts []string
	for _, part := range parts {
		if !strings.Contains(part, ":") {
			continue
		}
		key, value := utils.SplitKeyValue(part)

		switch key {
		case "org", "company":
			queryParts = append(queryParts, fmt.Sprintf("company=%s", value))
		case "hostname":
			domains := strings.Split(value, ".")
			if len(domains) >= 2 {
				queryParts = append(queryParts, fmt.Sprintf("domain=%s", strings.Join(domains[len(domains)-2:], ".")))
			}
		case "html":
			if strings.Contains(value, "@") {
				queryParts = append(queryParts, fmt.Sprintf("email=%s", value))
			}
		}
	}
	return strings.Join(queryParts, "&")
}
