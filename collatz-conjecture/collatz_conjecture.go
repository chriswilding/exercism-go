package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1
func CollatzConjecture(n int) (int, error) {
	var steps int

	if n <= 0 {
		return steps, errors.New("n must not be zero or negative")
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2

		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
