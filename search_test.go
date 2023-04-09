package grokking_test

import (
	"reflect"
	"testing"

	"github.com/bevzzz/grokking"
)

type SearchFunc = func([]string, string) (int, error)

var phoneBook = []string{
	"Alice",
	"Bob",
	"Clark",
	"Devlin",
	"Emmet",
	"Fatima",
	"Frankie",
	"Gigi",
}

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name   string
		search SearchFunc
		person string
		want   int
	}{
		{name: "simple search", search: grokking.SearchSimple, person: "Emmet", want: 4},
		{name: "binary search", search: grokking.SearchBinary, person: "Fatima", want: 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.search(phoneBook, tt.person)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n-got: %v\n+want: %v\n", got, tt.want)
			}
		})
	}
}
