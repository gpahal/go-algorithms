package numerical

// LCM return the lcm of a and b.
func LCM(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	return int((a * b) / GCD(a, b))
}
