package darts

import "math"

// Score calculates a darts score for given x and y coordinates.
func Score(x, y float64) int {
	radius := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case radius <= 1.0:
		return 10
	case radius <= 5.0:
		return 5
	case radius <= 10.0:
		return 1
	default:
		return 0
	}
}
