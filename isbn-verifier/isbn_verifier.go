package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	n := 10
	var sum int
	for _, r := range isbn {
		if n == 1 && r == 'X' {
			sum += n * 10
			n--
		}
		if unicode.IsDigit(r) {
			sum += n * int(r-'0')
			n--
		}
	}
	return n == 0 && sum%11 == 0
}
