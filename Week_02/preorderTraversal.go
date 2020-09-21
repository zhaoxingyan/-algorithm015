package Week_02

/*给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
1
\
2
/
3

输出: [1,2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//二叉树的前序遍历递归求解根---左子树---右子树

func preorderTraversal(root *TreeNode) []int {
	var res []int
	//terminator
	if root == nil {
		return nil
	}
	//process
	res = append(res, root.Val)
	//drill down
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

//迭代法实现前序遍历   根---左节点---右节点
//根节点压栈
//只要栈长度大于零弹出栈顶元素
//按照先进后出的规律，先压右节点后压左节点能保证出栈的时候按照左--右的顺序进行
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

