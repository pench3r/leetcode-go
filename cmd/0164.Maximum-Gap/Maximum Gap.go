package leetcode

import "fmt"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return 0
}

func quickSort(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	left := lo
	right := hi
	piovt := lo
	for lo < hi {
		for nums[lo] <= nums[piovt] && lo < hi {
			lo++
		}
		for nums[hi] >= nums[piovt] && lo < hi {
			hi--
		}
		if lo < hi {
			swap(nums, lo, hi)
		}
	}
	swap(nums, lo, piovt)
	quickSort(nums, left, lo-1)
	quickSort(nums, lo+1, right)
}

func swap(nums []int, l int, r int) {
	tmp := nums[l]
	nums[l] = nums[r]
	nums[r] = tmp
}
