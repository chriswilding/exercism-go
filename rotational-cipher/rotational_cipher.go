package rotationalcipher

import (
	"bytes"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var b bytes.Buffer

	for _, r := range plain {
		if unicode.IsLower(r) {
			b.WriteRune(rotate(r, shiftKey, 'z'))
		} else if unicode.IsUpper(r) {
			b.WriteRune(rotate(r, shiftKey, 'Z'))
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func rotate(r rune, shiftKey int, limit rune) rune {
	cipher := r + rune(shiftKey)
	if cipher > limit {
		cipher -= 26
	}
	return cipher
}
