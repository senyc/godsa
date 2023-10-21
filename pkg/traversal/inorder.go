package traversal

import (
	tree "github.com/senyc/godsa/internal/tree"
)

func Inorder(node *tree.TreeNode) []int {
	var l []int
	if node != nil {
		l = append(l, Inorder(node.Left)...)

		l = append(l, node.Val)

		l = append(l, Inorder(node.Right)...)
	}
	return l
}
