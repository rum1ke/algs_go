package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	tsetTable := []struct {
		sortedItems []int
		item        int
		must        int
	}{
		{sortedItems: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, item: 7, must: 7},
		{sortedItems: []int{333, 444, 555}, item: 111, must: -1},
		{sortedItems: []int{}, item: 78, must: -1},
	}

	for _, test := range tsetTable {
		res := BinarySearch(test.sortedItems, test.item)

		if res != test.must {
			t.Error("Result mismatch. Res, Must:", res, test.must)
		}
	}
}
