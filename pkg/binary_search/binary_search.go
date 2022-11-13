package binary_search

// Returns index of searched item in _sorted_ array.
// If requested item is not present in array it returns -1.
// Time Complexity: O(log n).
func BinarySearch(sortedItems []int, item int) int {
	var lowIdx int = 0
	var highIdx int = len(sortedItems) - 1
	var guess, midIdx int

	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		guess = sortedItems[midIdx]

		if guess == item {
			return midIdx
		}

		if guess > item {
			highIdx = midIdx - 1
		} else {
			lowIdx = midIdx + 1
		}
	}

	return -1
}
