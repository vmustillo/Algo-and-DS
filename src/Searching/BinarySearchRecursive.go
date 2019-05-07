package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	return binarySearchHelper(arr, 0, len(arr) - 1, target)
}

// Need sorted array
// Compare to target to the middle element
// If the target is larger, search the right half recursively from mid + 1
// If the target is larger, search the left half recursively from mid - 1
func binarySearchHelper(arr []int, left int, right int, target int) bool {
	if left > right {
		return false
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return true
	} else if target < arr[mid] {
		return binarySearchHelper(arr, 0, mid - 1, target)
	} else {
		return binarySearchHelper(arr, mid + 1, right, target)
	}
}

func main() {
	arr := []int{1, 2, 4, 5, 7, 9, 12, 18}
	target := 9
	if binarySearch(arr, target) {
		fmt.Println(target, "is in the array")
	} else {
		fmt.Println(target, "is not in the array")
	}
}