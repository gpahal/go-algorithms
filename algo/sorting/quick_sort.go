package sorting

// QuickSort sorts the slice in-place using the quick sort algorithm.
func QuickSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Choose the middle element as the pivot.
	lastIdx := arrLength - 1
	pivotIdx := arrLength >> 1
	pivotVal := arr[pivotIdx]

	arr[lastIdx], arr[pivotIdx] = pivotVal, arr[lastIdx]

	// Partition the slice.
	i := 0
	for j := 0; j < lastIdx; j++ {
		if arr[j] <= pivotVal {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	if i != lastIdx {
		arr[i], arr[lastIdx] = pivotVal, arr[i]
	}

	// Recursively sort the two partitions.
	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}
