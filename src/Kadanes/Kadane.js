// Find the maximum sub array given an array of integers
/**
 * This function returns the sum of the integers of the maximum contiguous subarray in value
 * The subarray must be contigous meaning all elements of the subarray must be consecutive
 * The way to calculate this is by using Kadane's algorithm
 *
 * @param {number[] nums}
 * @return {number}
 */
function findMaxSubArray(nums) {
    var length = nums.length; // precompute length of array so we do not have to more than once
    if (length === 0) {
        return 0;
    }
    else if (length === 1) {
        return nums[0];
    }
    var globalMax = nums[0], currentMax = nums[0]; // set global and current maximum subarrays to the first value of the array
    for (var i = 1; i < length; i++) { //currentMax is already the first element, so we can start at i = 1
        currentMax = Math.max(nums[i], currentMax + nums[i]); // find max between currentMax and currentMax + current element
        if (currentMax > globalMax) { // If this is greater than the maximum subarray, then this value is the new maximum subarray
            globalMax = currentMax;
        }
    }
    return globalMax;
}
var nums = [1, 6, -4, 2, 5, 3, -4, -1, 3];
console.log(findMaxSubArray(nums));
// Maximum subarray = 13
// 1 + 6 + (-4) + 2 + 5 + 3
