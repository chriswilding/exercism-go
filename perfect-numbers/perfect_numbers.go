package perfect

import "errors"

type Classification int

const (
	ClassificationPerfect   Classification = iota
	ClassificationAbundant  Classification = iota
	ClassificationDeficient Classification = iota
)

var ErrOnlyPositive error = errors.New("must be a positive integer")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}
	var sum int64
	limit := n/2 + 1
	for i := int64(1); i < limit; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
