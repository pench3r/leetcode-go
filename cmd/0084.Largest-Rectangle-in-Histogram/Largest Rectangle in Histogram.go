package leetcode


func largestRectangleArea(heights []int) int {
	stack, lenght, maxHeight := []int{}, 0, 0
	for inx := 0; inx <= len(heights); inx++ {
		height := 0
		if inx == len(heights) {
			height = 0
		} else {
			height = heights[inx]
		}
		if (len(stack) == 0) || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, inx)
		} else {
			stack_top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (len(stack) == 0) {
				lenght = inx
			} else {
				lenght = inx - 1 - stack[len(stack)-1]
			}
			maxHeight = max(maxHeight, lenght * heights[stack_top])
			inx--
		}
	}
	return maxHeight
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
