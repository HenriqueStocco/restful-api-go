package helpers

import (
	"strings"
)

func SetURL(method string, path string) string {
	const baseUrl string = "/api/v1"
	completePath := strings.Join([]string{baseUrl, path}, "")

	return strings.Join([]string{method, completePath}, " ")
}
