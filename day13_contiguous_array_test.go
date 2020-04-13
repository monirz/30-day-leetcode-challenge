package main

import "testing"

var testContiguousSubArray = []struct {
	nums     []int
	expected int
}{
	{[]int{0, 1}, 2},
	{[]int{0, 1, 0}, 2},
	{[]int{0, 0, 1, 0, 0, 0, 1, 1}, 6},
}

func TestFindMaxLength(t *testing.T) {

	for _, v := range testContiguousSubArray {

		result := findMaxLength(v.nums)

		if v.expected != result {
			t.Errorf("Expected %v got %v", v.expected, result)
		}
	}

}
