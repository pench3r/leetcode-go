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

func levelOrder1(root *structure.TreeNode) [][]int {
	var res [][]int
	dfs(root, -1, &res)
	return res
}

func dfs(root *structure.TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	curLevel := level + 1
	for len(*res) <= curLevel {
		*res = append(*res, []int{})
	}
	(*res)[curLevel] = append((*res)[curLevel], root.Val)
	dfs(root.Left, curLevel, res)
	dfs(root.Right, curLevel, res)
}
