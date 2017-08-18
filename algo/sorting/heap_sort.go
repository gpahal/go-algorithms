package sorting

// HeapSort sorts the slice in-place using the heap sort algorithm.
func HeapSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Build the heap.
	for i := arrLength/2 - 1; i >= 0; i -= 1 {
		heapify(arr, arrLength, i)
	}

	// Extract max values one by one, place them at the end and build the heap again.
	for i := arrLength - 1; i >= 0; i -= 1 {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, size, idx int) {
	largestIdx := idx
	leftIdx := idx<<1 + 1
	rightIdx := idx<<1 + 2

	if leftIdx < size && arr[leftIdx] > arr[largestIdx] {
		largestIdx = leftIdx
	}
	if rightIdx < size && arr[rightIdx] > arr[largestIdx] {
		largestIdx = rightIdx
	}

	if largestIdx != idx {
		arr[idx], arr[largestIdx] = arr[largestIdx], arr[idx]
		heapify(arr, size, largestIdx)
	}
}
