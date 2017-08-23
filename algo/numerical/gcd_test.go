package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/algo/numerical"
)

func TestGCD(t *testing.T) {
	cases := [][]int{
		{52, 78, 26},
		{8, -20, 4},
		{40, 40, 40},
		{0, 10, 10},
	}

	for _, c := range cases {
		gcd := numerical.GCD(c[0], c[1])
		if gcd != c[2] {
			t.Errorf("GCD: expected GCD to be %d, got %d", c[2], gcd)
		}
	}
}
