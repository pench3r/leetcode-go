package leetcode

import (
	"strconv"

	"github.com/pench3r/leetcode-go/structure"
)

func sumNumbers(root *structure.TreeNode) int {
	res, nums := 0, binaryTreeNums(root)
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		res += num
	}
	return res
}

func binaryTreeNums(root *structure.TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	tempLeft := binaryTreeNums(root.Left)
	for i := 0; i < len(tempLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+tempLeft[i])
	}
	tempRight := binaryTreeNums(root.Right)
	for i := 0; i < len(tempRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+tempRight[i])
	}
	return res
}
