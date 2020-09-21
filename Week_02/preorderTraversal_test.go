package Week_02

import (
	"fmt"
	"testing"
)

func TestPreOrderTraversal(t *testing.T) {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node3, nil}
	root := &TreeNode{1, nil, node2}
	fmt.Println(preorderTraversal(root))
}

func TestPreOrderTraversal2(t *testing.T) {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node3, nil}
	root := &TreeNode{1, nil, node2}
	fmt.Println(preorderTraversal2(root))
}
