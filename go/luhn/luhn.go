package luhn

import (
	"strconv"
	"strings"
)

func check_length(id string) bool {
	return len(id) > 1
}

func convert_to_slice_of_ints(id string) ([]int, error) {
	// error is returned if non-digit value is found
	ints := make([]int, len(id))

	for i, s := range strings.Split(id, "") {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = v
	}

	return ints, nil
}

func double_every_second_digit_and_return_sum(id_ints []int) int {
	// doubling starts from the right
	sum := 0

	id_ints_len := len(id_ints)

	for i := range id_ints {
		if i%2 != 0 && i != 0 {
			double := id_ints[id_ints_len-i-1] * 2
			if double > 9 {
				double -= 9
			}
			id_ints[id_ints_len-i-1] = double
		}
		sum += id_ints[id_ints_len-i-1]
	}
	return sum
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if !check_length(id) {
		return false
	}

	id_ints, err := convert_to_slice_of_ints(id)
	if err != nil {
		return false
	}

	id_ints_sum := double_every_second_digit_and_return_sum(id_ints)
	return id_ints_sum%10 == 0
}
