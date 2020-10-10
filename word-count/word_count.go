package wordcount

import (
	"regexp"
	"strings"
)

var find = regexp.MustCompile("\\b[a-zA-Z0-9]+[a-zA-Z0-9']*\\b")

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := make(Frequency)
	matches := find.FindAllString(phrase, -1)
	for _, match := range matches {
		lower := strings.ToLower(match)
		result[lower]++
	}
	return result
}
