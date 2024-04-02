package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	_len := len(a)
	if _len != len(b) {
		return 0, errors.New("The Hamming distance is only defined for sequences of equal length.")
	}

	distance := 0

	for i := 0; i < _len; i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
