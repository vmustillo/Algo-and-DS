package main

import (
	"fmt"
	"algo/searching/binarysearch"
)

func main() {
	arr := []int{1, 2, 4, 5, 7, 9, 12, 18}
	targets := []int{1, 18, 5, 9, 13}
	fmt.Println("Array: ", arr, "\nTargets: ", targets)

	for _, t := range targets {
		if binarysearch.Iterative(arr, t) {
			fmt.Println("\tIterative: ", t, "is in the array")
		} else {
			fmt.Println("\tIterative: ", t, "is not in the array")
		}
		if binarysearch.Recursive(arr, t) {
			fmt.Println("\tRecursive: ", t, "is in the array")
		} else {
			fmt.Println("\tRecursive: ", t, "is not in the array")
		}
	}
}