package leetcode

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	findCombinationSum3(1, k, n, c, &res)
	return res
}

func findCombinationSum3(index int, k int, target int, ans []int, res *[][]int) {
	if k == 0 && target == 0 {
		tmp := make([]int, len(ans))
		copy(tmp, ans)
		*res = append(*res, tmp)
		return
	}
	if k == 0 {
		return
	}
	for i := index; i <= 9; i++ {
		if target < i {
			break
		}
		ans = append(ans, i)
		findCombinationSum3(i+1, k-1, target-i, ans, res)
		ans = ans[:len(ans)-1]
	}
}
