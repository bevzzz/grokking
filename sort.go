package grokking

import (
	"math/rand"
)

// SortSelection sorts an slice of integers in a descending order using selection sort algorithm.
func SortSelection(in []int) (out []int, err error) {
	whichLargest := func(values []int) int {
		var idx int
		largest := values[idx]
		for i, v := range values[1:] {
			if v > largest {
				idx = i + 1
				largest = v
			}
		}
		return idx
	}

	for len(in) > 0 {
		idx := whichLargest(in)
		out = append(out, in[idx])

		upper := idx + 1
		if upper > len(in) {
			upper = len(in)
		}
		in = append(in[0:idx], in[upper:]...)
	}
	return out, nil
}

// Quicksort sorts an slice of integers in a descending order using quicksort algorithm.
func Quicksort(in []int) (out []int, err error) {
	// Base case
	if len(in) < 2 {
		return in, nil
	}

	// Optional special case, would also work without it.
	if len(in) == 2 {
		if in[0] > in[1] {
			return in, nil
		}
		return append(in[1:], in[0]), nil
	}

	// Recursive case
	pivotIdx := rand.Intn(len(in) - 1)
	pivot := in[pivotIdx]

	var smaller, larger []int
	for i, value := range in {
		if i == pivotIdx {
			continue
		}

		if value < pivot {
			smaller = append(smaller, value)
		} else {
			larger = append(larger, value)
		}
	}

	// Append larger + pivot + smaller for descending order
	out, err = Quicksort(larger)
	if err != nil {
		return in, err
	}
	out = append(out, pivot)

	sortedS, err := Quicksort(smaller)
	if err != nil {
		return in, err
	}
	return append(out, sortedS...), nil
}
