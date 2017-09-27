package numerical

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	return int((a * b) / GCD(a, b))
}

// LCMArray returns the least common multiple of all the elements in arr. If arr is nil or empty, 0 is returned.
func LCMArray(arr []int) int {
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

		ans = LCM(ans, arr[i])
	}

	return ans
}
