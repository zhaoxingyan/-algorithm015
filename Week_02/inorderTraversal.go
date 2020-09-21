package Week_02
/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
1
\
2
/
3

输出: [1,3,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
 */
//中序遍历  左子树---根---右子树
//1、递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

//用栈模拟
/*
每到一个节点A，根节点的访问在中间，将A入栈，遍历左子树
访问A---遍历右子树
访问完A后，A就可以出栈，因为和其左子树都已经访问完成
*/
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var res []int
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		res = append(res, top.Val)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			root = top.Right
		}
	}
	return res
}
