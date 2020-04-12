package main

import "testing"

var stoneTests = []struct {
	vals     []int
	expected int
}{
	{[]int{2, 7, 4, 1, 8}, 1},
	{[]int{1, 3}, 2},
}

func TestlastStoneWeight(t *testing.T) {

	for _, v := range stoneTests {
		result := lastStoneWeight(v.vals)

		if v.expected != result {
			t.Errorf("Expected %v got %v", v.expected, result)
		}
	}
}
