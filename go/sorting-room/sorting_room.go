package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	f := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", f)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if s, ok := fnb.(FancyNumber); ok {
		i, _ := strconv.Atoi(s.Value())
		return i
	} else {
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	f := float64(ExtractFancyNumber(fnb))
	return fmt.Sprintf("This is a fancy box containing the number %.1f", f)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var s string
	switch v := i.(type) {
	case int:
		s = DescribeNumber(float64(v))
	case float64:
		s = DescribeNumber(v)
	case NumberBox:
		s = DescribeNumberBox(v)
	case FancyNumberBox:
		s = DescribeFancyNumberBox(v)
	default:
		s = "Return to sender"
	}
	return s
}
