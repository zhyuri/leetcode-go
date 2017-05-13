package leetcode_go

import (
	"strconv"
)

// https://leetcode.com/problems/hamming-distance

func hammingDistance(x int, y int) int {
	bin := strconv.FormatInt(int64(x^y), 2)
	result := 0
	for _, b := range bin {
		if string(b) == "1" {
			result++
		}
	}
	return result
}
