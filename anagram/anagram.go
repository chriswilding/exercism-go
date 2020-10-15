package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	ls := strings.ToLower(subject)
	s := []byte(ls)
	sort.Slice(s, func(a, b int) bool {
		return s[a] < s[b]
	})
	var anagrams []string
	for _, candidate := range candidates {
		if l := len(subject); l != len(candidate) {
			continue
		}
		lc := strings.ToLower(candidate)
		if ls == lc {
			continue
		}
		c := []byte(lc)
		sort.Slice(c, func(a, b int) bool {
			return c[a] < c[b]
		})

		match := true

		for i := 0; i < len(c); i++ {
			if s[i] != c[i] {
				match = false
				break
			}
		}
		if match {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
