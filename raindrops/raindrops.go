package raindrops

import (
	"strconv"
)

// Convert converts a number into a string containing raindrop sounds
func Convert(n int) string {
	var output string

	if n%3 == 0 {
		output += "Pling"
	}
	if n%5 == 0 {
		output += "Plang"
	}
	if n%7 == 0 {
		output += "Plong"
	}

	if len(output) > 0 {
		return output
	}
	return strconv.Itoa(n)
}
