package fact

// Returns factorial of x.
func Fact(x int) int {
	if x == 1 {
		return x
	}

	return x * Fact(x-1)
}
