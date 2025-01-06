package mappings

var ShodanToCensys = map[string]string{
	"port":                "services.port",
	"protocol":            "services.service_name",
	"country":             "location.country_code",
	"city":                "location.city",
	"org":                 "autonomous_system.organization",
	"asn":                 "autonomous_system.asn",
	"product":             "services.software.product",
	"version":             "services.software.version",
	"http.status":         "services.http.response.status_code",
	"http.title":          "services.http.response.html_title",
	"http.server":         "services.http.response.headers.server",
	"ssl.version":         "services.tls.version",
	"ssl.cert.subject.cn": "services.tls.certificates.leaf_data.subject.common_name",
	"ssl.cert.issuer.cn":  "services.tls.certificates.leaf_data.issuer.common_name",
}

var CensysToShodan = map[string]string{
	"services.port":                                           "port",
	"services.service_name":                                   "protocol",
	"location.country_code":                                   "country",
	"location.city":                                           "city",
	"autonomous_system.organization":                          "org",
	"autonomous_system.asn":                                   "asn",
	"services.software.product":                               "product",
	"services.software.version":                               "version",
	"services.http.response.status_code":                      "http.status",
	"services.http.response.html_title":                       "http.title",
	"services.http.response.headers.server":                   "http.server",
	"services.tls.version":                                    "ssl.version",
	"services.tls.certificates.leaf_data.subject.common_name": "ssl.cert.subject.cn",
	"services.tls.certificates.leaf_data.issuer.common_name":  "ssl.cert.issuer.cn",
}
