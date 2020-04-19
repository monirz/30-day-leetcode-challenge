package main

import "testing"

var roatedSortedArrayTests = []struct {
	nums     []int
	target   int
	expected int
}{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
}

func TestSearchRotatedSortedArray(t *testing.T) {

	for _, v := range roatedSortedArrayTests {

		result := search(v.nums, v.target)

		if result != v.expected {
			t.Errorf("expected %v, got %v", v.expected, result)
		}
	}
}
