package utils

import "regexp"

func normalizeString(input string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return reg.ReplaceAllString(input, "")
}
