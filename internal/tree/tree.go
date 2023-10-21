package tree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func NewNode(val int) *TreeNode {
	r := new(TreeNode)
	r.Val = val
	return r
}

func NewTree(val int) *Tree {
	t := new(Tree)
	t.Root = NewNode(val)
	return t
}
