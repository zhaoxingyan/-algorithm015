package Week_03
/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

3
/ \
9  20
/  \
15   7

链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 前序遍历分为根|左子树|右子树
 中序遍历分为左子树|根|右子树
 前序遍历的第一个节点为根节点，由此可将中序遍历分为9|3|15，20，7    前序遍历为3|9|20，15，7
 以3为根节点的左子树为   前序遍历为9           中序遍历为9
            右子树     前序遍历为20，15，7   中序遍历为15，20，7
这是一个递归的过程，通过根节点在前序遍历中的pre_index对应的值可将中序遍历数组分为两个区间
根节点中序遍历下边i，为了方便将中序遍历的数据存放到hash表中，以节点值-节点索引的方式存放
左子树[left,i-1]
右子区间[i+1, right]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderHash := make(map[int]int)
	for k,v := range inorder {
		inorderHash[v] = k
	}
	return recursion(preorder, inorderHash, 0, 0, len(preorder)-1)
}

func recursion(preorder []int, inorderHash map[int]int, pre_root int ,left int, right int) *TreeNode {
	//递归终止条件（叶子节点）
	if left > right {
		return nil
	}
	val := preorder[pre_root]
	root := &TreeNode{val, nil, nil}
	i := inorderHash[val]//根节点在中序遍历中的索引
	Left := recursion(preorder, inorderHash, pre_root+1, left, i-1)

	Right := recursion(preorder, inorderHash, pre_root+i-left+1, i+1, right)
	root.Left = Left
	root.Right = Right
	return root
}
