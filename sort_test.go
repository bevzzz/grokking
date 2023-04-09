package grokking_test

import (
	"reflect"
	"testing"

	"github.com/bevzzz/grokking"
)

type SortFunc = func([]int) ([]int, error)

func TestSort(t *testing.T) {
	for _, tt := range []struct {
		name     string
		sort     SortFunc
		in, want []int
	}{
		{name: "selection", sort: grokking.SortSelection, in: []int{2, 1, 3, 8, 6}, want: []int{8, 6, 3, 2, 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sort(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n-got: %v\n+want: %v\n", got, tt.want)
			}
		})
	}
}
