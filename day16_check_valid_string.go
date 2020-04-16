package main

func checkValidString(s string) bool {

	leftCount := 0
	rightCount := 0
	l := len(s) - 1

	for i := 0; i <= l; i++ {

		if s[i] == '(' || s[i] == '*' {
			leftCount++
		} else {
			leftCount--
		}
		if s[l-i] == ')' || s[l-i] == '*' {
			rightCount++
		} else {
			rightCount--
		}

		if leftCount < 0 || rightCount < 0 {
			return false
		}

	}

	return true

}
