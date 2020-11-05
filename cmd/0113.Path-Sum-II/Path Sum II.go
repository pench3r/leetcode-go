package leetcode

import "github.com/pench3r/leetcode-go/structure"

func pathSum(root *structure.TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findpath(root, sum, []int(nil), slice)
	return slice
}

func findpath(root *structure.TreeNode, sum int, stack []int, slice [][]int) [][]int {
	if root == nil {
		return slice
	}
	sum -= root.Val
	stack = append(stack, root.Val)
	if root.Left == nil && root.Right == nil && sum == 0 {
		slice = append(slice, append([]int{}, stack...))
	}
	slice = findpath(root.Left, sum, stack, slice)
	slice = findpath(root.Right, sum, stack, slice)
	return slice
}
