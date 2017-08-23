package numerical

// GCD return the greatest common divisor of a and b.
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
