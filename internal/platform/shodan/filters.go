package shodan

var Filters = map[string]string{
	// Network
	"ip":       `ip:"%s"`,
	"port":     `port:%s`,
	"net":      `net:"%s"`,
	"asn":      `asn:"%s"`,
	"hostname": `hostname:"%s"`,
	"cidr":     `net:"%s"`,

	// Geographic
	"country":  `country:"%s"`,
	"city":     `city:"%s"`,
	"region":   `region:"%s"`,
	"postal":   `postal:"%s"`,
	"geo":      `geo:"%s"`,
	"location": `location:"%s"`,

	// Organizations
	"org":      `org:"%s"`,
	"isp":      `isp:"%s"`,
	"customer": `customer:"%s"`,

	// Devices & Products
	"product":      `product:"%s"`,
	"version":      `version:"%s"`,
	"devicetype":   `devicetype:"%s"`,
	"os":           `os:"%s"`,
	"brand":        `brand:"%s"`,
	"model":        `model:"%s"`,
	"manufacturer": `manufacturer:"%s"`,

	// Web Technologies
	"server":     `server:"%s"`,
	"powered-by": `powered-by:"%s"`,
	"framework":  `framework:"%s"`,
	"cms":        `cms:"%s"`,
	"platform":   `platform:"%s"`,

	// HTTP Specific
	"http.component":   `http.component:"%s"`,
	"http.status":      `http.status:%s`,
	"http.title":       `http.title:"%s"`,
	"http.server":      `http.server:"%s"`,
	"http.html":        `http.html:"%s"`,
	"http.location":    `http.location:"%s"`,
	"http.host":        `http.host:"%s"`,
	"http.auth_scheme": `http.auth_scheme:"%s"`,
	"http.waf":         `http.waf:"%s"`,
	"http.cors":        `http.cors:"%s"`,
	"http.headers":     `http.headers:"%s"`,

	// SSL/TLS
	"ssl":                  `ssl:%s`,
	"ssl.version":          `ssl.version:"%s"`,
	"ssl.cert.subject.cn":  `ssl.cert.subject.cn:"%s"`,
	"ssl.cert.issuer.cn":   `ssl.cert.issuer.cn:"%s"`,
	"ssl.cert.fingerprint": `ssl.cert.fingerprint:"%s"`,
	"ssl.cert.serial":      `ssl.cert.serial:"%s"`,
	"ssl.cert.expired":     `ssl.cert.expired:%s`,
	"ssl.cipher":           `ssl.cipher:"%s"`,
	"ssl.alpn":             `ssl.alpn:"%s"`,
	"ssl.jarm":             `ssl.jarm:"%s"`,

	// Vulnerabilities
	"vuln":          `vuln:"%s"`,
	"vuln.verified": `vuln.verified:%s`,
	"vuln.cvss":     `vuln.cvss:%s`,
	"vuln.cve":      `vuln.cve:"%s"`,

	// Cloud & Infrastructure
	"cloud.provider":  `cloud.provider:"%s"`,
	"cloud.service":   `cloud.service:"%s"`,
	"cloud.region":    `cloud.region:"%s"`,
	"container":       `container:"%s"`,
	"container.name":  `container.name:"%s"`,
	"container.image": `container.image:"%s"`,

	// IoT & Industrial
	"has_screenshot":  `has_screenshot:%s`,
	"screenshot.hash": `screenshot.hash:"%s"`,
	"industrial":      `industrial:%s`,
	"ics":             `ics:"%s"`,
	"scada":           `scada:"%s"`,
	"plc":             `plc:"%s"`,

	// Authentication & Access
	"auth":     `auth:"%s"`,
	"password": `password:%s`,
	"proxy":    `proxy:%s`,
	"vpn":      `vpn:%s`,

	// Metadata
	"timestamp": `timestamp:"%s"`,
	"tag":       `tag:"%s"`,
	"hash":      `hash:"%s"`,
	"comment":   `comment:"%s"`,

	// Database Specific
	"database":   `database:"%s"`,
	"db.name":    `db.name:"%s"`,
	"db.version": `db.version:"%s"`,

	// Advanced Filters
	"all":             `all:"%s"`,
	"bitcoin.address": `bitcoin.address:"%s"`,
	"blockchain":      `blockchain:"%s"`,
	"cryptocurrency":  `cryptocurrency:"%s"`,
	"favicon.hash":    `favicon.hash:"%s"`,
	"html_hash":       `html_hash:"%s"`,
	"malware":         `malware:"%s"`,
}
