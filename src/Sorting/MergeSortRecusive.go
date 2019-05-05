package main

import "fmt"

// mergeSort first divides the array down into arrays with single values
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	return merge(mergeSort(arr[0:middle]), mergeSort(arr[middle:]))
}

// merge takes two arrays as input and sorts them into a new array.
// Runs through both arrays comparingelements and inserting the smaller into the new array
// Once one array has been iterated through, if there are any elements left in the other array, we know those elements are already sorted and can be appended to the end of the resultant array 
func merge(leftHalf []int, rightHalf []int) []int {
	left := 0
	right := 0
	var mergedArr []int

	for left < len(leftHalf) && right < len(rightHalf) {
		if leftHalf[left] <= rightHalf[right] {
			mergedArr = append(mergedArr, leftHalf[left])
			left++
		} else {
			mergedArr = append(mergedArr, rightHalf[right])
			right++
		}
	}

	if left < len(leftHalf) {
		mergedArr = append(mergedArr, leftHalf[left:]...)
	}

	if right < len(rightHalf) {
		mergedArr = append(mergedArr, rightHalf[right:]...)
	}
	
	return mergedArr
}

func main() {
	arr := []int{51, 4, 8, 29, 54, 12, 73, 43}
	fmt.Println("Unsorted array: ", arr)
	newArr := mergeSort(arr)
	fmt.Println("Sorted array: ", newArr)
}