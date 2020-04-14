package main

import "testing"

var stringShiftTests = []struct {
	s        string
	shift    [][]int
	expected string
}{
	{"abc", [][]int{{0, 1}, {1, 2}}, "cab"},
	{"abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}, "efgabcd"},
	{"wpdhhcj", [][]int{{0, 7}, {1, 7}, {1, 0}, {1, 3}, {0, 3}, {0, 6}, {1, 2}}, "hcjwpdh"},
}

func TestStringShift(t *testing.T) {

	for _, v := range stringShiftTests {

		result := stringShift(v.s, v.shift)

		if result != v.expected {
			t.Errorf("expected %v got %v", v.expected, result)
		}
	}
}
