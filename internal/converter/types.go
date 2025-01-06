package converter

type QueryParams map[string]interface{}

type Converter interface {
	ToCensys(query string) string
	ToShodan(query string) string
	ToHunter(query string) string
}
