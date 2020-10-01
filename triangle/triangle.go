package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	for _, side := range sides {
		if side <= 0 || math.IsInf(side, 0) || math.IsNaN(side) {
			return NaT
		}
	}

	sort.Float64s(sides)
	if sides[0]+sides[1] < sides[2] {
		return NaT
	}

	if a == b && b == c && c == a {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
