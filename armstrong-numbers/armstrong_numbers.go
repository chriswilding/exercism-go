package armstrong

import (
	"math"
)

func IsNumber(input int) bool {
	length := math.Floor(math.Log10(float64(input))) + 1

	var sum int
	rem := input

	for rem > 0 {
		nth := float64(rem % 10)
		rem /= 10
		sum += int(math.Pow(nth, length))
	}

	return input == sum
}
