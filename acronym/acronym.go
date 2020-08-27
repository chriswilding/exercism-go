package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	delimiter := true
	var b strings.Builder

	for _, r := range s {
		if delimiter && unicode.IsLetter(r) {
			b.WriteRune(unicode.ToUpper(r))
			delimiter = false
		} else if unicode.IsSpace(r) || r == '-' {
			delimiter = true
		}
	}

	return b.String()
}
