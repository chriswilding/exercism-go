package scale

import (
	"strings"
)

var flat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharp = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var intervals = map[rune]int{'m': 1, 'M': 2, 'A': 3}

func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	var chromatic []string
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chromatic = flat
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		chromatic = sharp
	default:
		chromatic = sharp
	}

	var chromaticScaleIndex int
	chromaticScaleLookup := strings.Title(tonic)

	for i, candidate := range chromatic {
		if candidate == chromaticScaleLookup {
			chromaticScaleIndex = i
			break
		}
	}

	output := make([]string, 0, len(interval))

	for _, il := range interval {
		output = append(output, chromatic[chromaticScaleIndex])
		in := intervals[il]
		chromaticScaleIndex = chromaticScaleIndex + in
		if chromaticScaleIndex >= len(chromatic) {
			chromaticScaleIndex -= len(chromatic)
		}
	}

	return output
}
