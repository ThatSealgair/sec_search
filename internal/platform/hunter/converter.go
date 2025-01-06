package hunter

import (
	"fmt"
	"net/url"
	"strings"
)

type QueryConverter struct{}

func (c *QueryConverter) ToHunter(query string) string {
	// Converting to self, just return the query
	return query
}

func (c *QueryConverter) ToShodan(query string) string {
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}

	var parts []string

	if company := values.Get("company"); company != "" {
		parts = append(parts, fmt.Sprintf(`org:"%s"`, company))
	}
	if domain := values.Get("domain"); domain != "" {
		parts = append(parts, fmt.Sprintf(`hostname:"%s"`, domain))
	}
	if email := values.Get("email"); email != "" {
		parts = append(parts, fmt.Sprintf(`html:"%s"`, email))
	}

	return strings.Join(parts, " ")
}

func (c *QueryConverter) ToCensys(query string) string {
	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}

	var parts []string

	if company := values.Get("company"); company != "" {
		parts = append(parts, fmt.Sprintf(`autonomous_system.organization: "%s"`, company))
	}
	if domain := values.Get("domain"); domain != "" {
		parts = append(parts, fmt.Sprintf(`services.http.response.html: "*%s*"`, domain))
	}
	if email := values.Get("email"); email != "" {
		parts = append(parts, fmt.Sprintf(`services.http.response.html: "*%s*"`, email))
	}

	return strings.Join(parts, " AND ")
}
