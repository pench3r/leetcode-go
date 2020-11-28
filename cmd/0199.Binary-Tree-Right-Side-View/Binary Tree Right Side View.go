package leetcode

import "github.com/pench3r/leetcode-go/structure"

func rightSideView(root *structure.TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*structure.TreeNode{}
	curLevel, nextLevel := 1, 0
	queue = append(queue, root)
	for len(queue) > 0 {
		if curLevel > 0 {
			node := queue[0]
			if curLevel == 1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevel++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevel++
			}
			queue = queue[1:]
			curLevel--
		}
		if curLevel == 0 {
			curLevel = nextLevel
			nextLevel = 0
		}
	}
	return res
}
