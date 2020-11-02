package leetcode

import (
	"math"

	"github.com/pench3r/leetcode-go/structure"
)

func isValidBST(root *structure.TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}

func isValidbst(root *structure.TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v > min && v < max && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}

func isValidBST1(root *structure.TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *structure.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
