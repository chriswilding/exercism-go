package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	output := []rune(plain)

	for i, r := range output {
		if unicode.IsLower(r) {
			output[i] = rotate(r, shiftKey, 'z')
		} else if unicode.IsUpper(r) {
			output[i] = rotate(r, shiftKey, 'Z')
		}
	}

	return string(output)
}

func rotate(r rune, shiftKey int, limit rune) rune {
	cipher := r + rune(shiftKey)
	if cipher > limit {
		cipher -= 26
	}
	return cipher
}
