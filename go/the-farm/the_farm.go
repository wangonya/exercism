package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows int
	msg  string
}

func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fodder, err := f.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	factor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return (factor * fodder) / float64(cows), nil
}

func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(f, cows)
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}
	return nil
}
