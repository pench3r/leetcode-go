package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(candidates []int, target int, index int, ans []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			tmp := make([]int, len(ans))
			copy(tmp, ans)
			*res = append(*res, tmp)
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		ans = append(ans, candidates[i])
		findCombinationSum(candidates, target-candidates[i], i, ans, res)
		ans = ans[:len(ans)-1]
	}
}
