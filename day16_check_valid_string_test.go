package main

import "testing"

var checkValidStringTests = []struct {
	s        string
	expected bool
}{
	{"()", true},
	{"(", false},
	{")", false},
	{"())", false},
	{"(())", true},
	{"(*))", true},
}

func TestCheckValidString(t *testing.T) {

	for _, v := range checkValidStringTests {

		result := checkValidString(v.s)

		if result != v.expected {

			t.Errorf("expected %v got %v", v.expected, result)
		}
	}

}
