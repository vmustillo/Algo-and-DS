package main

import (
	"fmt"
)

/*
	Given n (size of array) and s (sum to add up to),
	find a coniguous subarray in which the elements add up to the sum.
	If there is no subarray return -1.
	Ex: n = 5, s = 12
	arr = [1, 2, 3, 7, 5]
	output: [2, 3, 7]
*/

func main() {
	nums := []int{1, 2, 3, 7, 5}
	sum := 12
	arr := contigSubarrayEqToSum(nums, sum)
	fmt.Println(arr)

}

func contigSubarrayEqToSum(nums []int, sum int) []int {
	start := 0
	end := 0

	if len(nums) == 0 {
		return []int{-1}
	}

	contiguousSum := nums[end]
	var contigSubarray []int
	contigSubarray = append(contigSubarray, nums[end])

	for end < len(nums) {
		fmt.Println(contigSubarray)
		if contiguousSum == sum {
			return contigSubarray
		} else if contiguousSum < sum {
			end++
			if end < len(nums) {
				contigSubarray = append(contigSubarray, nums[end])
				contiguousSum += nums[end]
			}
		} else if contiguousSum > sum && start < end {
			start++
			contiguousSum -= contigSubarray[0]
			contigSubarray = contigSubarray[1:]
		}
	}

	return []int{-1}
}
