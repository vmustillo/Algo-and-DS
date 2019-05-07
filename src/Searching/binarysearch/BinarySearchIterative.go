package binarysearch

// Need sorted array
// Compare to target to the middle element
// If the target is larger, search the right half recursively from mid + 1
// If the target is larger, search the left half recursively from mid - 1
func Iterative(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return true
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return false
}