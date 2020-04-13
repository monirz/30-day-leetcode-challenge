package main

func findMaxLength(nums []int) int {
	// [0,1]

	hm := make(map[int]int)
	sum := 0
	maxLen := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	for i := 0; i < len(nums); i++ {

		sum += nums[i]

		if sum == 0 {
			maxLen = i + 1
		}

		if _, ok := hm[sum]; ok {

			if maxLen < i-hm[sum] {
				maxLen = i - hm[sum]
			}
		} else {
			hm[sum] = i
		}

	}

	return maxLen

}
