package main

import "fmt"

func main() {

	nums := []int{1, 3}

	// [1,2,2, 4,7 ]
	// [1,2,2,3]
	//[1,1,2]
	//
	res := lastStoneWeight(nums)

	fmt.Println(res)
}
