package leetcode

import (
	"github.com/pench3r/leetcode-go/structure"
)

func isSameTree(p *structure.TreeNode, q *structure.TreeNode) bool {
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
