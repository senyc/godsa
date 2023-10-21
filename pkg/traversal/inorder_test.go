package traversal_test

import (
	"testing"
	tree "github.com/senyc/godsa/internal/tree"
	"github.com/senyc/godsa/pkg/traversal"
	"reflect"
)

func TestInorder(t *testing.T) {
	tr := tree.NewTree(5)
	tr.Root.Left = tree.NewNode(4)
	tr.Root.Right = tree.NewNode(6)
	tr.Root.Right.Left = tree.NewNode(7)
	tr.Root.Right.Right = tree.NewNode(8)
	tr.Root.Left.Left = tree.NewNode(2)
	tr.Root.Left.Right = tree.NewNode(3)
	tr.Root.Left.Left.Left = tree.NewNode(1)

	solution := []int{1,2,4,3,5,7,6,8}
	if result := traversal.Inorder(tr.Root); ! reflect.DeepEqual(result, solution) {
		t.Error(result, "does not equal", solution)
	}
}
