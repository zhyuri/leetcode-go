package leetcode_go

// https://leetcode.com/problems/frog-jump

// TODO: still meet "Memory Limit Exceed" error
func canCross(stones []int) bool {
	stoneMap := make(map[int][]int)
	stoneMap[0] = []int{1}

	for i := 1; i < len(stones); i++ {
		stoneMap[stones[i]] = []int{}

	}
	for i := 0; i < len(stones)-1; i++ {
		stone := stones[i]
		for _, step := range stoneMap[stone] {
			gap := step + stone
			if gap == stones[len(stones)-1] {
				return true
			}
			if _, ok := stoneMap[gap]; ok {
				stoneMap[gap] = append(stoneMap[gap], step)
				if step-1 > 0 {
					stoneMap[gap] = append(stoneMap[gap], step-1)
				}
				stoneMap[gap] = append(stoneMap[gap], step+1)
			}

		}

	}

	return false
}
