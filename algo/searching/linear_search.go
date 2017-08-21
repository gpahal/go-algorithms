package searching

// LinearSearch linearly searches the slice for the first occurrence of key and returns the index. If the key is not
// found, -1 is returned.
func LinearSearch(arr []int, key int) int {
	for i, el := range arr {
		if el == key {
			return i
		}
	}

	return -1
}
