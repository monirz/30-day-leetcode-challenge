package main

import "testing"

var minPathSumTests = []struct {
	grid     [][]int
	expected int
}{
	{
		[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},

		7},
}

func TestMinPathsSum(t *testing.T) {

	for _, v := range minPathSumTests {

		result := minPathSum(v.grid)

		if result != v.expected {
			t.Errorf("expected %v got %v", v.expected, result)
		}
	}
}
