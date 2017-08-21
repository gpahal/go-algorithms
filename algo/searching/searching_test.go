package searching_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gpahal/go-algorithms/algo/searching"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandSlice(length int) []int {
	if length <= 0 {
		return []int{}
	}

	arr := make([]int, length)
	for i := 0; i < length; i += 1 {
		arr[i] = (rand.Int() % 50) + 1 // random number in the range 1..50
	}

	return arr
}

func checkSearchResult(arr []int, key, idx int) (bool, int) {
	if len(arr) < idx+1 {
		for i, el := range arr {
			if el == key {
				return false, i
			}
		}

		return false, -1
	}

	if idx >= 0 {
		if arr[idx] != key {
			return false, -1
		}

		for i := 0; i < idx; i++ {
			if arr[i] == key {
				return false, i
			}
		}
	} else {
		for i, el := range arr {
			if el == key {
				return false, i
			}
		}
	}

	return true, idx
}

func assertSearchFn(t *testing.T, name string, fn func([]int, int) int) bool {
	for i := 0; i < 10; i += 1 {
		length := (rand.Int() % 20) + 1
		arr := generateRandSlice(length)

		var key int
		if i == 5 {
			key = -1
		} else if i%2 == 0 {
			key = arr[rand.Int()%length]
		} else {
			key = (rand.Int() % 50) + 1
		}

		idx := fn(arr, key)
		correct, correctIdx := checkSearchResult(arr, key, idx)
		if !correct {
			t.Errorf("%s %v for %d: expected %d, got %d", name, arr, key, correctIdx, idx)
			return false
		}
	}

	return true
}

func TestLinearSearch(t *testing.T) {
	assertSearchFn(t, "LinearSearch", searching.LinearSearch)
}
