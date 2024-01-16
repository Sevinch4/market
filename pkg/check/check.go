package check

import (
	"unicode"
)

func PhoneNumber(phone string) bool {
	for _, p := range phone {
		if p == '+' {
			continue
		} else if !unicode.IsNumber(p) {
			return false
		}
	}
	return true
}
