package Week_03

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor (t *testing.T) {
	node4 := &TreeNode{4,nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node2 := &TreeNode{2, node7, node4}
	node6 := &TreeNode{6, nil, nil}
	node0 := &TreeNode{0, nil, nil}
	node8 := &TreeNode{8, nil, nil}
	node1 := &TreeNode{1, node0, node8}
	node5 := &TreeNode{5, node6, node2}
	root := &TreeNode{3, node5, node1}

	ancestor := lowestCommonAncestor(root, node6,node8)
	ancestor2 := lowestCommonAncestor(root, node0,node8)
	fmt.Println(ancestor.Val)
	fmt.Println(ancestor2.Val)
}
