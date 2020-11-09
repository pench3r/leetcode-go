package leetcode

import (
	"fmt"

	"github.com/pench3r/leetcode-go/structure"
)

// recursive
func recoverTree(root *structure.TreeNode) {
	var prev, target1, target2 *structure.TreeNode
	_, target1, target2 = inOrderTraversal(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}

func inOrderTraversal(root, prev, target1, target2 *structure.TreeNode) (*structure.TreeNode, *structure.TreeNode, *structure.TreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = inOrderTraversal(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root
	prev, target1, target2 = inOrderTraversal(root.Right, prev, target1, target2)
	return prev, target1, target2
}

// morris
func morrisTravelsal(root *structure.TreeNode) {
	for root != nil {
		if root.Left != nil {
			tmp := root.Left
			for (tmp.Right != nil) && (tmp.Right != root) {
				tmp = tmp.Right
			}
			if tmp.Right == nil {
				tmp.Right = root
				root = root.Left
			}
			if tmp.Right == root {
				tmp.Right = nil
				// visite root val
				fmt.Println(root.Val)
				root = root.Right
			}
		} else {
			// visit root val
			fmt.Println(root.Val)
			root = root.Right
		}
	}
}
