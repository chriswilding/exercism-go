package strand

import "strings"

var mapping = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna strings.Builder
	for _, base := range dna {
		compliment := mapping[base]
		rna.WriteRune(compliment)
	}
	return rna.String()
}
