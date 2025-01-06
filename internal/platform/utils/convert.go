package utils

import (
	"fmt"
	"strings"
)

func SplitQuery(query, delimiter string) []string {
	var parts []string
	var currentPart strings.Builder
	var inQuotes bool
	var quoteChar rune

	for i, char := range query {
		switch {
		case char == '"' || char == '\'':
			if !inQuotes {
				inQuotes = true
				quoteChar = char
				currentPart.WriteRune(char)
			} else if quoteChar == char {
				inQuotes = false
				currentPart.WriteRune(char)

				if i == len(query)-1 || query[i+1] == ' ' {
					parts = append(parts, currentPart.String())
					currentPart.Reset()
				}
			} else {
				currentPart.WriteRune(char)
			}
		case char == ' ' && !inQuotes:
			if currentPart.Len() > 0 {
				parts = append(parts, currentPart.String())
				currentPart.Reset()
			}
		default:
			currentPart.WriteRune(char)
		}
	}

	if currentPart.Len() > 0 {
		parts = append(parts, currentPart.String())
	}

	return parts
}

func SplitKeyValue(part string) (string, string) {
	parts := strings.SplitN(part, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}
	key := strings.TrimSpace(parts[0])
	value := strings.Trim(strings.TrimSpace(parts[1]), `"`)
	return key, value
}

func FormatValue(value string, needsQuotes bool) string {
	if needsQuotes {
		return fmt.Sprintf(`"%s"`, value)
	}
	return value
}
