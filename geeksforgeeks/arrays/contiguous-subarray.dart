/*
	Given n (size of array) and s (sum to add up to),
	find a coniguous subarray in which the elements add up to the sum.
	If there is no subarray return -1.
	Ex: n = 5, s = 12
	arr = [1, 2, 3, 7, 5]
	output: [2, 3, 7]
*/

main() {
	List nums = [1, 2, 3, 7, 5];
	int sum = 12;
	var arr = contigSubarrayEqToSum(nums, sum);
  print(arr);
}

List contigSubarrayEqToSum(List nums, int sum) {
	int start = 0;
	int end = 0;

	if(nums.length == 0) {
		return [-1];
	}

	int contiguousSum = nums[end];
	List contigSubarray = [nums[end]];

  while(end < nums.length) {
    if(contiguousSum == sum) {
      return contigSubarray;
    } else if(contiguousSum < sum) {
      end++;
      if(end < nums.length) {
        contigSubarray.add(nums[end]);
        contiguousSum += nums[end];
      }
    } else if(contiguousSum > sum && start < end) {
      contiguousSum -= contigSubarray[0];
      contigSubarray.removeAt(0);
    }
  }

	return [-1];
}
