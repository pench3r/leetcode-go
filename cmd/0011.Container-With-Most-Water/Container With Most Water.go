package leetcode

func maxArea(height []int) int {
	res, left, right, hgt := 0, 0, len(height)-1, 0
	for left < right {
		if (height[left] > height[right]) {
			hgt = height[right]
			right--
		} else {
			hgt = height[left]
			left++
		}
		res = max(res, (right - left + 1) * hgt)
	}
	return res
}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}