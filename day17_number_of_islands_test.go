package main

import "testing"

var numIslandsTests = []struct {
	grid     [][]byte
	expected int
}{
	{
		[][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		},
		1},
	{
		[][]byte{
			[]byte("11000"),
			[]byte("11000"),
			[]byte("00100"),
			[]byte("00011"),
		},
		3},
}

func TestNumIslands(t *testing.T) {

	for _, v := range numIslandsTests {
		result := numIslands(v.grid)

		if result != v.expected {
			t.Errorf("expected %v got %v", v.expected, result)
		}

	}
}
