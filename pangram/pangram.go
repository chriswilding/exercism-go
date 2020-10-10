package pangram

import "unicode"

func IsPangram(phrase string) bool {
	if phrase == "" {
		return false
	}

	seen := make(map[rune]struct{}, 26)
	for _, r := range phrase {
		if l := unicode.ToLower(r); unicode.IsLetter(l) {
			seen[l] = struct{}{}
		}
	}
	return len(seen) == 26
}
