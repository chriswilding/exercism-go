package strand

import "strings"

var mapping = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var b strings.Builder
	for _, base := range dna {
		compliment := mapping[base]
		b.WriteRune(compliment)
	}
	return b.String()
}
