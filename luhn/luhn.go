package luhn

import (
	"unicode"
)

func Valid(input string) bool {
	runes := []rune(input)
	var digits int
	var double bool
	var sum int

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		digits++
		v := int(r - '0')
		if double {
			v *= 2
		}
		if v > 9 {
			v -= 9
		}
		sum += v
		double = !double
	}
	return digits > 1 && sum%10 == 0
}
