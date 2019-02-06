// Find the maximum sub array given an array of integers
// [1, 2, 6, -3, 2, -7, 5, 6, 4]

/**
 * This function returns the sum of the integers of the maximum contiguous subarray in value
 * The subarray must be contigous meaning all elements of the subarray must be consecutive
 * The way to calculate this is by using Kadane's algorithm
 * 
 * @param {number[] nums}
 * @return {number}
 */
function findMaxSubArray(nums: number[]) {
    var length: number = nums.length; // precompute length of array so we do not have to more than once
    if(length === 0) {
        return 0;
    } else if(length === 1) {
        return nums[0];
    }
    var globalMax: number = nums[0], currentMax: number = nums[0]; // set global and current maximum subarrays to the first value of the array
    for(let i: number = 1; i < length; i++) { //currentMax is already the first element, so we can start at i = 1
        currentMax = Math.max(currentMax, currentMax + nums[i]); // find max between currentMax and currentMax + current element
        if(currentMax > globalMax) { // If this is greater than the maximum subarray, then this value is the new maximum subarray
            globalMax = currentMax;
        }
    }
    return globalMax; 
}