package Week_03

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

链接：https://leetcode-cn.com/problems/combinations
*/

func combine(n int, k int) [][]int {
	if n < 0 || n < k {
		return nil
	}
	var path []int
	var res  [][]int
	dfs(n, k, 1, path, &res)
	return res
}

func dfs(n int, k int, index int, path []int, res *[][]int) {
	if len(path) == k {
		tmpPath := make([]int, k)
		copy(tmpPath, path)
		*res = append(*res, tmpPath)
		return
	}
	for i := index; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, res)
		path = path[:len(path)-1]
	}
}
