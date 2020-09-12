package scale

import (
	"strings"
)

var flat = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharp = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var intervals = map[rune]int{'m': 1, 'M': 2, 'A': 3}

func findScale(tonic string) []string {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return flat
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		return sharp
	default:
		return sharp
	}
}

func findIndexInScale(scale []string, tonic string) int {
	for i, candidate := range scale {
		if candidate == tonic {
			return i
		}
	}
	return 0
}

func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	chromatic := findScale(tonic)
	chromaticScaleIndex := findIndexInScale(chromatic, strings.Title(tonic))

	output := make([]string, 0, len(interval))

	for _, i := range interval {
		output = append(output, chromatic[chromaticScaleIndex])
		n := intervals[i]
		chromaticScaleIndex += n
		if chromaticScaleIndex >= len(chromatic) {
			chromaticScaleIndex -= len(chromatic)
		}
	}

	return output
}
