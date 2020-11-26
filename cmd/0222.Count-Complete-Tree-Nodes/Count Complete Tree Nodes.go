package leetcode

import "github.com/pench3r/leetcode-go/structure"

func countNodes(root *structure.TreeNode) int {
	if root == nil {
		return 0
	}
	res, queue := 1, []*structure.TreeNode{}
	curLevel, nextLevel := 1, 0
	queue = append(queue, root)
	for len(queue) > 0 {
		if curLevel > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevel++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevel++
			}
			curLevel--
		}
		if curLevel == 0 {
			res += nextLevel
			curLevel = nextLevel
			nextLevel = 0
		}

	}
	return res
}
