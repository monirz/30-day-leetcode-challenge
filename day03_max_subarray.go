package main

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	currentSum := maxSum

	for i := 1; i < len(nums); i++ {

		currentSum = max(nums[i]+currentSum, nums[i])

		maxSum = max(maxSum, currentSum)

	}

	return maxSum

}
