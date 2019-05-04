package main

import "fmt"

// MergeSort will be used to half the array and 
// func mergeSort(arr []int) {
// 	mergeHelper(arr, 0, len(arr) - 1)
// }

// //
// func mergeHelper(arr []int, left int, right int) {
// 	if left < right {
// 		middle := (left + right) / 2
// 		mergeHelper(arr, left, middle)
// 		mergeHelper(arr, middle + 1, right)
// 		mergeHalves(arr, left, right)
// 	}
// }

// // Merge is
// func mergeHalves(arr []int, left int, right int) {
// 	leftEnd := (left + right) / 2
// 	rightStart := leftEnd + 1
// 	temp := arr
// 	i := left

// 	for left <= leftEnd && rightStart <= right {
// 		fmt.Println(temp, left, rightStart)
// 		if arr[left] <= arr[rightStart] {
// 			temp[i] = arr[left]
// 			left++
// 		} else {
// 			temp[i] = arr[rightStart]
// 			rightStart++
// 		}
// 		i++
// 	}

// 	for left <= leftEnd {
// 		temp[i] = arr[left]
// 		left++
// 		i++
// 	}

// 	for rightStart <= right {
// 		temp[i] = arr[rightStart]
// 		rightStart++
// 		i++
// 	}
	
// 	arr = temp
// }

func mergeSort(arr []int) []int {

	if(len(arr) == 1) {
		return arr
	}
	middle := len(arr) / 2

	return merge(mergeSort(arr[0:middle]), mergeSort(arr[middle:]))
}

func merge(leftHalf []int, rightHalf []int) []int {
	left := 0
	right := 0
	leftEnd := len(leftHalf)
	rightEnd := len(rightHalf)
	fmt.Println("leftEnd: ", leftEnd)
	fmt.Println("leftEnd: ", rightEnd)
	mergedArr := make([]int, leftEnd + rightEnd - 1)
	i := 0

	for left < leftEnd && right < rightEnd {
		fmt.Println(mergedArr)
		if leftHalf[left] <= rightHalf[right] {
			mergedArr[i] = leftHalf[left]
			left++
		} else {
			mergedArr[i] = rightHalf[right]
			right++
		}
		i++
	}
	
	mergedArr = append(mergedArr, leftHalf[left:]...)
	mergedArr = append(mergedArr, rightHalf[right:]...)

	return mergedArr
}

func main() {
	arr := []int{51, 4, 8, 29, 54, 12, 73, 43}
	//arr := []int{51, 4, 8, 29}
	fmt.Println("Unsorted array: ", arr)
	newArr := mergeSort(arr)
	fmt.Println("Sorted array: ", newArr)
}