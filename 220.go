package leetcode_go

// https://leetcode.com/problems/contains-duplicate-iii

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)

	for i := 0; i <= length/2+1; i++ {
		for j := i + 1; j-i <= k && j < length; j++ {
			if nums[i]-nums[j] >= 0 && nums[i]-nums[j] <= t {
				return true
			}
			if nums[j]-nums[i] >= 0 && nums[j]-nums[i] <= t {
				return true
			}
		}
	}

	return false
}
