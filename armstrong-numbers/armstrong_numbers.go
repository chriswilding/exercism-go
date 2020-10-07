package armstrong

import "math"

func IsNumber(input int) bool {
	length := int(math.Floor(math.Log10(float64(input)))) + 1

	var sum int
	for i := 0; i < length; i++ {
		nth := input % int(math.Pow10(i+1)) / int(math.Pow10(i))
		sum += int(math.Pow(float64(nth), float64(length)))
	}

	return input == sum
}
