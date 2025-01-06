package censys

var Filters = map[string]string{
	// Basic Host Information
	"ip":                             `ip: %s`,
	"ip_prefix":                      `ip_prefix: %s`,
	"autonomous_system.asn":          `autonomous_system.asn: %s`,
	"autonomous_system.name":         `autonomous_system.name: "%s"`,
	"autonomous_system.organization": `autonomous_system.organization: "%s"`,

	// Location Information
	"location.country":      `location.country: "%s"`,
	"location.country_code": `location.country_code: "%s"`,
	"location.city":         `location.city: "%s"`,
	"location.province":     `location.province: "%s"`,
	"location.postal_code":  `location.postal_code: "%s"`,

	// Services
	"services.port":             `services.port: %s`,
	"services.service_name":     `services.service_name: "%s"`,
	"services.software.product": `services.software.product: "%s"`,
	"services.software.version": `services.software.version: "%s"`,

	// HTTP/HTTPS
	"services.http.response.status_code":    `services.http.response.status_code: %s`,
	"services.http.response.html_title":     `services.http.response.html_title: "%s"`,
	"services.http.response.headers.server": `services.http.response.headers.server: "%s"`,

	// TLS/SSL
	"services.tls.version": `services.tls.version: "%s"`,
	"services.tls.certificates.leaf_data.subject.common_name": `services.tls.certificates.leaf_data.subject.common_name: "%s"`,
	"services.tls.certificates.leaf_data.issuer.common_name":  `services.tls.certificates.leaf_data.issuer.common_name: "%s"`,
	"services.tls.certificates.leaf_data.serial_number":       `services.tls.certificates.leaf_data.serial_number: "%s"`,
}
