package main

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	//modified binary serach
	for left < right {

		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	start := left //swap the value
	left = 0
	right = len(nums) - 1 //reset the value

	if target >= nums[start] && target <= nums[right] {
		left = start
	} else {
		right = start
	}

	//pure binary search
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {

			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
