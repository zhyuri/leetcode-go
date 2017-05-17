package leetcode_go

import (
	"sort"
)

// https://leetcode.com/problems/anagrams

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for _, str := range strs {
		sortStr := sortString(str)
		strMap[sortStr] = append(strMap[sortStr], str)
	}

	var result [][]string
	for _, str := range strMap {
		if len(str) <= 0 {
			continue
		}
		result = append(result, str)
	}

	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
