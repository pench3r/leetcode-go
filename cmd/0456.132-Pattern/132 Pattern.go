package leetcode

import "math"

func find132pattern(nums []int) bool {
	num3, stack :=  math.MinInt64, []int{}
	if len(nums) < 3 {
		return false
	}
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] < num3 {
			return true
		}
		for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
			num3 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}