package main

func isHappy(n int) bool {

	hash := make(map[int]bool)

	for n != 1 {

		current := n
		sum := 0

		for current != 0 {
			sum += (current % 10) * (current % 10)
			current = current / 10
		}

		if _, ok := hash[sum]; ok {
			return false
		}

		hash[n] = true
		n = sum

	}

	return true

}
