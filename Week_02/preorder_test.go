package Week_02

import (
	"fmt"
	"testing"
)


func TestPreOrder(t *testing.T) {
	node5 := &Node{5, nil}
	node6 := &Node{6, nil}
	node4 := &Node{4, nil}
	node2 := &Node{2, nil}
	node3Children := []*Node{node5, node6}
	node3 := &Node{3, node3Children}
	rootChildren := []*Node{node3, node2, node4}
	root := &Node{1, rootChildren}
	fmt.Println(preorder(root))
}

func TestPreOrder2(t *testing.T) {
	node5 := &Node{5, nil}
	node6 := &Node{6, nil}
	node4 := &Node{4, nil}
	node2 := &Node{2, nil}
	node3Children := []*Node{node5, node6}
	node3 := &Node{3, node3Children}
	rootChildren := []*Node{node3, node2, node4}
	root := &Node{1, rootChildren}
	fmt.Println(preorder2(root))
}
