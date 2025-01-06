package hunter

var Filters = map[string]string{
	// Domain Search
	"domain":           `domain="%s"`,
	"company":          `company="%s"`,
	"company_industry": `company_industry="%s"`,
	"company_size":     `company_size="%s"`,

	// Email Finder
	"first_name": `first_name="%s"`,
	"last_name":  `last_name="%s"`,
	"full_name":  `full_name="%s"`,
	"email":      `email="%s"`,

	// Professional Info
	"position":     `position="%s"`,
	"seniority":    `seniority="%s"`,
	"department":   `department="%s"`,
	"linkedin_url": `linkedin_url="%s"`,
	"phone_number": `phone_number="%s"`,
	"twitter":      `twitter="%s"`,

	// Email Pattern
	"pattern":    `pattern="%s"`,
	"email_type": `email_type="%s"`,
	"status":     `status="%s"`,
	"type":       `type="%s"`,

	// Verification
	"disposable": `disposable=%s`,
	"webmail":    `webmail=%s`,
	"accept_all": `accept_all=%s`,
	"score":      `score=%s`,

	// Additional Parameters
	"limit":               `limit=%s`,
	"offset":              `offset=%s`,
	"max_results":         `max_results=%s`,
	"required_attributes": `required_attributes="%s"`,

	// Boolean Filters
	"personal_email": `personal_email=%s`,
	"generic_email":  `generic_email=%s`,
	"active":         `active=%s`,
	"verified":       `verified=%s`,
}
