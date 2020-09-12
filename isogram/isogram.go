package isogram

import "unicode"

// IsIsogram returns true if a given word or phrase does not repeat a letter
func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		l := unicode.ToLower(r)
		if _, ok := seen[l]; ok {
			return false
		}
		seen[l] = struct{}{}
	}
	return true
}
