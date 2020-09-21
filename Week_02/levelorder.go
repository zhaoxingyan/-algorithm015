package Week_02
/*
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

例如，给定一个 3叉树 :
返回其层序遍历:

[
     [1],
     [3,2,4],
     [5,6]
]
 

说明:

树的深度不会超过 1000。
树的节点总数不会超过 5000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
*/
/*
层序遍历使用队列模拟实现遍历,队列特点先进先出
1、根节点加入队列，队列为空时停止迭代
2、遍历该层次前使用一个size（counter）计数器来记录该层中的数据个数，计数器从零到counter时该层停止遍历
3、把同层的节点全部移除队列【确保每次迭代前队列中的数据为同层数据】把队列头部的counter数的节点移除，该层节点如果有子节点加入队列，等待下次输出
*/
func levelorder(root *Node) [][]int {
	var res [][]int
	var queue []*Node
	if root == nil {
		return res
	}
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		res = append(res, []int{})
		counter := len(queue)
		for i := 0; i < counter; i++ {
			res[level] = append(res[level], queue[i].Val)
			if queue[i].Children != nil {
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[counter:]
		level++
	}
	return res
}
