package leetcode_go

import (
	"math"
)

// https://leetcode.com/problems/contains-duplicate-ii

func containsNearbyDuplicate(nums []int, k int) bool {
	integers := make(map[int]int)
	for i, num := range nums {
		if j, ok := integers[num]; ok {
			if math.Abs(float64(i - j)) <= float64(k) {
				return true
			}
		}
		integers[num] = i
	}

	return false
}
