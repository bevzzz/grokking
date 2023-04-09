package grokking

import (
	"errors"
)

// SearchSimple returns the index of the value in the lookup
// using simple search algorithm. Make sure lookup is alphabetically sorted.
func SearchSimple(lookup []string, value string) (int, error) {
	for i, check := range lookup {
		if value == check {
			return i, nil
		}
	}
	return 0, errors.New("value not present in array")
}

// SearchSimple returns the index of the value in the lookup
// using binary search algorithm. Make sure lookup is alphabetically sorted.
func SearchBinary(lookup []string, value string) (int, error) {
	if len(lookup) == 0 {
		return 0, errors.New("lookup is empty")
	}

	var inUpper bool
	start, end := 0, len(lookup)
	split := start + int((end-start)/2)

	// Split lookup in two, check to which half our value belongs
	if my, check := value[0], lookup[split][0]; my == check && lookup[split] == value {
		return split, nil // lucky
	} else if inUpper = my > check; inUpper {
		lookup = lookup[split+1:]
	} else {
		lookup = lookup[:split]
	}

	// Recursive call with a smaller lookup range
	idx, err := SearchBinary(lookup, value)
	if err != nil {
		return 0, err
	}
	if inUpper {
		idx += split + 1
	}
	return idx, nil
}
