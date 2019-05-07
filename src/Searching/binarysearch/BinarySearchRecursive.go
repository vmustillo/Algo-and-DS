package binarysearch

func Recursive(arr []int, target int) bool {
	return RecursiveHelper(arr, 0, len(arr) - 1, target)
}

// Need sorted array
// Compare to target to the middle element
// If the target is larger, search the right half recursively from mid + 1
// If the target is larger, search the left half recursively from mid - 1
func RecursiveHelper(arr []int, left int, right int, target int) bool {
	if left > right {
		return false
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return true
	} else if target < arr[mid] {
		return RecursiveHelper(arr, 0, mid - 1, target)
	} else {
		return RecursiveHelper(arr, mid + 1, right, target)
	}
}