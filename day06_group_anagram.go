package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	anagrams := [][]string{}

	m := make(map[string][]string)

	for _, v := range strs {
		s := strings.Split(v, "")
		sort.Strings(s)
		sorted := strings.Join(s, "")

		if _, ok := m[sorted]; !ok {

			newArr := []string{}
			m[sorted] = newArr
		}

		m[sorted] = append(m[sorted], v)
	}

	for _, v := range m {
		anagrams = append(anagrams, v)
	}

	return anagrams
}
