package leetcode

import (
	"github.com/pench3r/leetcode-go/structure"
)

func levelOrder(root *structure.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*structure.TreeNode{}
	queue = append(queue, root)
	curCounter, nextLevel, tmp, res := 1, 0, []int{}, [][]int{}
	for len(queue) > 0 {
		node := queue[0]
		if curCounter > 0 {
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevel++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevel++
			}
			curCounter--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curCounter == 0 {
			curCounter = nextLevel
			res = append(res, tmp)
			nextLevel = 0
			tmp = []int{}
		}
	}
	return res
}