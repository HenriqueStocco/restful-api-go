package helpers

import "strings"

func SetURL(path string) string {
	const baseUrl string = "/api/v1"

	return strings.Join([]string{baseUrl, path}, "")
}
