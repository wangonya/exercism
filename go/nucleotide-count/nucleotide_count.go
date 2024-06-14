package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]uint8

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for i := 0; i < len(d); i++ {
		switch d[i] {
		case 'A':
			h['A'] += 1
		case 'C':
			h['C'] += 1
		case 'G':
			h['G'] += 1
		case 'T':
			h['T'] += 1
		default:
			return h, errors.New("Invalid")
		}
	}
	return h, nil
}
