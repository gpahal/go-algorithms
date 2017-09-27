package numerical_test

import (
	"testing"

	"github.com/gpahal/go-algorithms/algo/numerical"
)

func TestLCM(t *testing.T) {
	cases := [][3]int{
		{52, 78, 156},
		{8, -20, 40},
		{1, 3, 3},
		{40, 40, 40},
		{0, 10, 0},
	}

	for _, c := range cases {
		lcm := numerical.LCM(c[0], c[1])
		if lcm != c[2] {
			t.Errorf("LCM %d, %d: expected LCM to be %d, got %d", c[0], c[1], c[2], lcm)
		}
	}
}
