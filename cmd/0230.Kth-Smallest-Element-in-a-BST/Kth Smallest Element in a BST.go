package leetcode

import (
	"github.com/pench3r/leetcode-go/structure"
)

func kthSmallest(root *structure.TreeNode, k int) int {
	ans, count := 0, 0
	inorderTraversal(root, k, &count, &ans)
	return ans
}

func inorderTraversal(root *structure.TreeNode, k int, count *int, ans *int) {
	if root != nil {
		inorderTraversal(root.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = root.Val
		}
		inorderTraversal(root.Right, k, count, ans)
	}
}
