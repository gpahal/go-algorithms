package numerical

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// GCDArray returns the greatest common divisor of all the elements in arr. If arr is nil or empty, 0 is returned.
func GCDArray(arr []int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	if arrLen == 1 {
		return arr[0]
	}

	ans := arr[0]
	for i := 1; i < arrLen; i++ {
		if ans == 0 {
			return 0
		}

		ans = GCD(ans, arr[i])
	}

	return ans
}
