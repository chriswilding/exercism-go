package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number into a string containing raindrop sounds
func Convert(n int) string {
	var b strings.Builder

	if n%3 == 0 {
		b.WriteString("Pling")
	}
	if n%5 == 0 {
		b.WriteString("Plang")
	}
	if n%7 == 0 {
		b.WriteString("Plong")
	}

	if b.Len() > 0 {
		return b.String()
	}
	return strconv.Itoa(n)
}
