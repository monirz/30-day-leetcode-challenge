package main

import "testing"

var productExceptSelfTests = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	// {[]}
}

func TestProductExceptSelf(t *testing.T) {

	for _, v := range productExceptSelfTests {

		result := productExceptSelf(v.nums)

		for i, val := range result {
			if val != v.expected[i] {
				t.Errorf("expected %v got %v", v.expected, result)
			}
		}

	}

}
