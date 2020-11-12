package utils

import "github.com/microcosm-cc/bluemonday"

func SanitizeInput(sanitizer *bluemonday.Policy, arr ...*string) {
	for index, val := range arr {
		*arr[index] = sanitizer.Sanitize(*val)
	}
}
