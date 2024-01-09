package utils

import "strings"

func GenerateSlug(title string) string {
	slug := strings.ReplaceAll(normalizeString(title), " ", "-")
	return strings.ToLower(slug)
}
