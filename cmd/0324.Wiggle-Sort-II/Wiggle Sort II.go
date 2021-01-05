package leetcode

import "sort"

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	array := make([]int, len(nums))
	copy(array, nums)
	sort.Ints(array)
	left := (len(nums)+1)/2 - 1
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			nums[i] = array[right]
			right--
		} else {
			nums[i] = array[left]
			left--
		}
	}
}
