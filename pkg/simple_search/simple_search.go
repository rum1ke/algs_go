package simple_search

// Finds and returns first met index of searched item in array.
// If requested item is not present in array it returns -1.
// Time Complexity: O(n).
func SimpleSearch(list []int, item int) int {
	for idx, guess := range list {
		if guess == item {
			return idx
		}
	}

	return -1
}
