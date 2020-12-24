package utils

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
)

func SanitizeInput(sanitizer *bluemonday.Policy, arr ...*string) {
	fmt.Println("Sanitizing input")
	for index, val := range arr {
		*arr[index] = sanitizer.Sanitize(*val)
	}
}
