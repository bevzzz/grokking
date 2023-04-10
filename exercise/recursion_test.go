package exercise_test

import (
	"testing"

	"github.com/bevzzz/grokking/exercise"
)

func TestSummaries(t *testing.T) {
	for _, tt := range []struct {
		name string
		f    func(...int) int
		in   []int
		want int
	}{
		{"Sum: add 2 numbers", exercise.Sum, []int{5, 3, 1}, 9},
		{"Count: number of elements", exercise.Count, []int{5, 3, 1, 2, 2}, 5},
		{"Max: find the largest number", exercise.Max, []int{0, 354, 122}, 354},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f(tt.in...); got != tt.want {
				t.Errorf("\n-got: %v\n+want: %v\n", got, tt.want)
			}
		})
	}
}
