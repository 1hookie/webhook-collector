package masker

import "strings"

func MaskEmail(email string) string {
	if email == "" {
		return ""
	}

	mask := strings.Map(func(r rune) rune {
		if r != '@' {
			return '*'
		}
		return r
	}, email)

	return mask
}
