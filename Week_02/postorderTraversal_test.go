package Week_02

import (
	"fmt"
	"testing"
)

func TestPostOrderTraveral(t *testing.T)  {
	node6 := &TreeNode{6,nil, nil}
	node5 := &TreeNode{5,nil, nil}
	node4 := &TreeNode{4,nil, nil}
	node3 := &TreeNode{3, node6, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
	fmt.Println(postorderTraversal(root))
}

func TestPostOrderTraveral2(t *testing.T)  {
	node6 := &TreeNode{6,nil, nil}
	node5 := &TreeNode{5,nil, nil}
	node4 := &TreeNode{4,nil, nil}
	node3 := &TreeNode{3, node6, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
	fmt.Println(postorderTraversal2(root))
}

