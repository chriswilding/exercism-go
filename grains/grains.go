package grains

import "errors"

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}

	return uint64(1) << (n - 1), nil
}

func Total() uint64 {
	var sum uint64

	for i := 1; i <= 64; i++ {
		result, _ := Square(i)
		sum += result
	}

	return sum
}
