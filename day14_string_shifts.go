package main

func stringShift(s string, shift [][]int) string {

	leftShift := 0
	rightShift := 0
	var result string

	for i := 0; i < len(shift); i++ {

		if shift[i][0] == 0 {
			leftShift = leftShift + shift[i][1]
		} else {
			rightShift = rightShift + shift[i][1]
		}
	}

	if leftShift > rightShift {
		leftShift -= rightShift
		leftShift = leftShift % len(s)

		result += s[leftShift:]
		result += s[:leftShift]

	} else {
		rightShift -= leftShift
		rightShift = rightShift % len(s)

		result += s[len(s)-rightShift:]
		result += s[:len(s)-rightShift]
	}

	return result
}
