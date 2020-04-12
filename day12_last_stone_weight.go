package main

import (
	"sort"
)

//using sort with each turn
func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		l := len(stones) - 1

		if stones[l] == stones[l-1] {
			stones = stones[:l-1]
		} else {

			x := stones[l] - stones[l-1]

			stones = stones[:l-1]
			stones = append(stones, x)

		}

	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}
