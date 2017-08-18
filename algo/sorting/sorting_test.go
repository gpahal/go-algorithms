package sorting_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/gpahal/go-algorithms/algo/sorting"
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
		arr[i] = rand.Int()
	}

	return arr
}

func assertSortFn(t *testing.T, name string, fn func([]int)) bool {
	for i := 0; i < 5; i += 1 {
		original := generateRandSlice(20)
		sorted := make([]int, 20)
		copy(sorted, original)
		fn(sorted)

		for j := 0; j < 19; j += 1 {
			if sorted[j] > sorted[j+1] {
				t.Errorf("%s %v: got %v (%d > %d)", name, original, sorted, sorted[j], sorted[j+1])
				return false
			}
		}
	}

	return true
}

func TestQuickSort(t *testing.T) {
	assertSortFn(t, "QuickSort", sorting.QuickSort)
}

func TestMergeSort(t *testing.T) {
	assertSortFn(t, "MergeSort", sorting.MergeSort)
}

func TestHeapSort(t *testing.T) {
	assertSortFn(t, "HeapSort", sorting.HeapSort)
}

func benchmarkSortFn(b *testing.B, fn func([]int), length int) {
	arr := generateRandSlice(length)
	for i := 0; i < b.N; i++ {
		fn(arr)
	}
}

func BenchmarkQuickSort_10000(b *testing.B) {
	benchmarkSortFn(b, sorting.QuickSort, 10000)
}

func BenchmarkMergeSort_10000(b *testing.B) {
	benchmarkSortFn(b, sorting.MergeSort, 10000)
}

func BenchmarkHeapSort_10000(b *testing.B) {
	benchmarkSortFn(b, sorting.HeapSort, 10000)
}