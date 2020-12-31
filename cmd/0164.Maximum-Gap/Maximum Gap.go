package leetcode

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	quickSort(nums, 0, len(nums)-1)
	res := 0
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i]-nums[i-1])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func quickSort(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	left := lo
	right := hi
	piovt := lo
	for lo < hi {
		for nums[hi] >= nums[piovt] && lo < hi {
			hi--
		}
		for nums[lo] <= nums[piovt] && lo < hi {
			lo++
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

func maximumGap1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		maxNum = max(maxNum, nums[i])
	}
	exp := 1
	aux := make([]int, len(nums))
	for (maxNum / exp) > 0 {
		count := make([]int, 10)
		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}
		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
			count[(nums[i]/exp)%10] = tmp
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}
	res := 0
	for i := 1; i < len(aux); i++ {
		res = max(res, aux[i]-aux[i-1])
	}
	return res
}
