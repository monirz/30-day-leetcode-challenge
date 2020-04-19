package main

func moveZeroes(nums []int) {

	pos := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]

			pos++
		}

	}
}
