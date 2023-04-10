package exercise

import "reflect"

// Sum calculates the total of the passed values using recursion.
func Sum(values ...int) int {
	switch len(values) {
	case 0:
		return 0
	case 1:
		return values[0]
	default:
		return values[0] + Sum(values[1:]...)
	}
}

func Count(values ...int) int {
	if reflect.DeepEqual(values, []int{}) {
		return 0
	}
	return 1 + Count(values[1:]...)
}

func Max(values ...int) int {
	if len(values) > 2 {
		return Max(values[0], Max(values[1:]...))
	}

	if values[0] > values[1] {
		return values[0]
	}
	return values[1]
}
