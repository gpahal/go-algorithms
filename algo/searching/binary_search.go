package searching

// BinarySearch searches the sorted slice for the key using the binary search algorithm. If the key is not
// found, -1 is returned.
func BinarySearch(arr []int, key int) int {
	return binarySearchHelper(arr, key, 0)
}

func binarySearchHelper(arr []int, key int, idxToBeAdded int) int {
	arrLength := len(arr)
	if arrLength == 0 {
		return -1
	}

	mid := arrLength >> 1
	midVal := arr[mid]
	if midVal == key {
		return mid + idxToBeAdded
	} else if midVal > key {
		return binarySearchHelper(arr[:mid], key, idxToBeAdded)
	} else {
		return binarySearchHelper(arr[mid+1:], key, idxToBeAdded+mid+1)
	}
}
