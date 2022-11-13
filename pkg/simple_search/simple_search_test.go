package simple_search

import "testing"

func TestSimpleSearch(t *testing.T) {
	testTable := []struct {
		items []int
		item  int
		must  int
	}{
		{items: []int{3, 3, 1, 5, 8, 3, 0, 2}, item: 3, must: 0},
		{items: []int{123, 88, 1616}, item: 55, must: -1},
		{items: []int{}, item: 1, must: -1},
	}

	for _, test := range testTable {
		res := SimpleSearch(test.items, test.item)

		if res != test.must {
			t.Error("Result mismatch. Res, Must:", res, test.must)
		}
	}
}
