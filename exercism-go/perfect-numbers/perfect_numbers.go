package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = -1
	ClassificationPerfect   Classification = 0
	ClassificationAbundant  Classification = 1
)

var ErrOnlyPositive error = errors.New("ErrOnlyPositive")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}

	var sum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	switch {
	case sum < n:
		return ClassificationDeficient, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}
