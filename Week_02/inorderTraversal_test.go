package Week_02

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T)  {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node3, nil}
	root := &TreeNode{1, nil, node2}
	fmt.Println(inorderTraversal(root))
}

func TestInorderTraversal2(t *testing.T) {
	node6 := &TreeNode{6,nil, nil}
	node5 := &TreeNode{5,nil, nil}
	node4 := &TreeNode{4,nil, nil}
	node3 := &TreeNode{3, node6, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
	fmt.Println(inorderTraversal2(root))

}
