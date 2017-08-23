package numerical

// Fibonacci return the nth fibonacci number. NOTE: Fibonacci(0) = 0 and Fibonacci(1) = 1.
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	pprev := 0
	prev := 1
	for i := n - 1; i > 0; i-- {
		pprev, prev = prev, pprev+prev
	}

	return prev
}
