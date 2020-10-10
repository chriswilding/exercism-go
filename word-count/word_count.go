package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	lower := strings.ToLower(phrase)
	matches := strings.FieldsFunc(lower, f)
	result := make(Frequency)
	for _, match := range matches {
		trimmed := strings.Trim(match, "'")
		result[trimmed]++
	}
	return result
}

func f(r rune) bool {
	switch {
	case unicode.IsDigit(r):
		return false
	case unicode.IsLetter(r):
		return false
	case r == '\'':
		return false
	}
	return true
}
