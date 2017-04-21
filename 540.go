package leetcode_go

// https://leetcode.com/problems/single-element-in-a-sorted-array
// Accepted

func singleNonDuplicate(nums []int) int {
	for i := 1; i < len(nums)-1; i = i + 2 {
		if nums[i] != nums[i-1] {
			if nums[i-1] == -1 {
				return nums[i]
			} else {
				return nums[i-1]
			}
		} else {
			nums[i] = -1
			nums[i-1] = -1
		}
	}
	return nums[len(nums)-1]
}
