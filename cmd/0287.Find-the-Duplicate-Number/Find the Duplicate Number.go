package leetcode

func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
