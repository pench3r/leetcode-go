package leetcode

import (
	"github.com/pench3r/leetcode-go/structure"
)

func isSymmetric(root *structure.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(p *structure.TreeNode, q *structure.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}
