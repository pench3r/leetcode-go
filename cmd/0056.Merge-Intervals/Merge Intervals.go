package leetcode

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		return res
	}
	quickSort(intervals, 0, len(intervals)-1)
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] < intervals[i][0] {
			res = append(res, prev)
			prev = intervals[i]
		} else {
			prev[1] = max(intervals[i][1], prev[1])
		}
	}
	res = append(res, prev)
	return res
}

func quickSort(nums [][]int, start int, end int) {
	if start >= end {
		return
	}
	left := start
	right := end
	piovt := nums[start]
	for left < right {
		for (nums[right][0] >= piovt[0]) && left < right {
			right--
		}
		for (nums[left][0] <= piovt[0]) && left < right {
			left++
		}
		if left < right {
			swap(nums, left, right)
		}
	}
	swap(nums, start, left)
	quickSort(nums, start, left-1)
	quickSort(nums, right+1, end)
}

func swap(nums [][]int, m int, n int) {
	tmp := nums[m]
	nums[m] = nums[n]
	nums[n] = tmp
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
