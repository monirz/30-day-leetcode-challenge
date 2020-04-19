package main

func countElements(arr []int) int {

	m := make(map[int]bool)

	count := 0

	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
	}

	for _, v := range arr {
		if _, ok := m[v+1]; ok {
			count++
		}
	}

	return count
}
