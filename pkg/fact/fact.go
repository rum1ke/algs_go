package fact

// Returns factorial of x.
// Time Complexity: O(n).
func Fact(x int) int {
	if x == 1 {
		return x
	}

	return x * Fact(x-1)
}
