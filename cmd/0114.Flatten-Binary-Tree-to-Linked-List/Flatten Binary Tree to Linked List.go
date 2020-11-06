package leetcode

import "github.com/pench3r/leetcode-go/structure"

// 递归
func flatten(root *structure.TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	curRight := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = curRight
}

// 非递归
func flatten1(root *structure.TreeNode) {
	list, cur := []int{}, &structure.TreeNode{}
	preorder(root, &list)
	cur = root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &structure.TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}
}

func preorder(root *structure.TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}
