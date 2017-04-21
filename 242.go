package leetcode_go

// https://leetcode.com/problems/valid-anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[int32]int)

	for _, c := range s {
		charMap[c]++
	}
	for _, c := range t {
		charMap[c]--
	}
	for _, m := range charMap {
		if m != 0 {
			return false
		}
	}
	return true
}
