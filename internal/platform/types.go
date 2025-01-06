package platform

type Platform string

const (
	Shodan Platform = "shodan"
	Censys Platform = "censys"
	Hunter Platform = "hunter"
)

func (p Platform) String() string {
	return string(p)
}

type Converter interface {
	ToCensys(query string) string
	ToShodan(query string) string
	ToHunter(query string) string
}
