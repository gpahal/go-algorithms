package sorting

// MergeSort sorts the slice in-place using the merge sort algorithm.
func MergeSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 2 {
		if arrLength == 2 && arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return
	}

	// Recursively sort the two halves of the slice.
	mid := arrLength >> 1
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])

	merge(arr, mid)
}

func merge(arr []int, mid int) {
	// Check if merge is required.
	if arr[mid-1] <= arr[mid] {
		return
	}

	arrLength := len(arr)

	// Temporary slice to remember the first half of the original slice.
	tmpArr := make([]int, mid)
	copy(tmpArr, arr[:mid])

	// Merge the two parts.
	idx, left, right := 0, 0, mid
	for {
		if tmpArr[left] <= arr[right] {
			arr[idx] = tmpArr[left]
			left++
			if left == mid {
				break
			}
		} else {
			arr[idx] = arr[right]
			right++
			if right == arrLength {
				copy(arr[idx+1:], tmpArr[left:mid])
				break
			}
		}

		idx++
	}
}
