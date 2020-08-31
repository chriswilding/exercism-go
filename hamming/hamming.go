package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	var distance int
	if len(a) != len(b) {
		return distance, errors.New("strands must have the same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
