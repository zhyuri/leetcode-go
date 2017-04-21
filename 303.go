package leetcode_go

// https://leetcode.com/problems/range-sum-query-immutable
// Better Solution: https://leetcode.com/articles/range-sum-query-immutable/

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	result := 0
	for index, num := range this.Nums {
		if i <= index && index <= j {
			result += num
		}
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
