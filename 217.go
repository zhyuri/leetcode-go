package leetcode_go

// https://leetcode.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	integers := make(map[int]int)

	for _, num := range nums {
		if _, ok := integers[num]; ok {
			return true
		}
		integers[num] = 1
	}

	return false
}
