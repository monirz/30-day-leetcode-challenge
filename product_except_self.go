package main

func productExceptSelf(nums []int) []int {

	l := len(nums)

	leftProduct := make([]int, l)
	rightProduct := make([]int, l)

	result := make([]int, l)

	leftProduct[0] = 1

	for i := 1; i < l; i++ {

		leftProduct[i] = nums[i-1] * leftProduct[i-1]

	}

	rightProduct[l-1] = 1

	for i := l - 2; i >= 0; i-- {
		rightProduct[i] = nums[i+1] * rightProduct[i+1]
	}

	for i := 0; i < l; i++ {
		result[i] = leftProduct[i] * rightProduct[i]
	}
	return result
}
