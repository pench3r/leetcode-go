package leetcode;

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for inx, num := range nums {
		if p, ok := m[target-num]; ok {
			return []int{p, inx}
		}
		m[num] = inx
	}
	return nil
}

