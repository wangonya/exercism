package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n < 1 {
		return count, errors.New("negative")
	}

	for n != 1 {
		count++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return count, nil
}
