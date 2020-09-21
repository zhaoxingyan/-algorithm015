package Week_02
/*
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :

返回其前序遍历: [1,3,5,6,2,4]。
前序遍历的特点：
对于一个树，访问顺序为根--左孩子--右孩子
*/
/*
两种解法：
1、递归
2、借助栈实现前序遍历
*/
type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	//terminator
	if root == nil {
		return nil
	}
	//process
	res = append(res, root.Val)

	//drill down
	for i := 0; i < len(root.Children); i++ {
		res = append(res, preorder(root.Children[i])...)
	}
	return res
}
//迭代方式实现前序遍历
//1、根节点压栈
//2、迭代弹栈【只要栈中有节点就执行弹栈操作】
////2.1 弹出根节点加入结果集中
////2.2 如果有子节点将子节点逆序压栈【保障最左边的子节点最先被弹出】迭代直到栈为空
func preorder2(root *Node) []int {
	var stack []*Node  //栈
	var res []int      //结果集
	if root == nil {
		return nil     //空树直接返回
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, top.Val)
		//子节点逆序压栈
		for i := len(top.Children)-1; i >=0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return res
}
