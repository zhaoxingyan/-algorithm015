package Week_03

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

链接：https://leetcode-cn.com/problems/permutations
*/

func permute(nums []int) [][]int {
	//回溯算法
	lenth := len(nums)    //数组长度
	res := [][]int{}    //保存结果集--二位切片
	path := []int{}     //保存中间结果--以为切片
	used := make([]bool, lenth)    //保存数组中的数据是否被加入中间结果
	depth := 0           //遍历的层级
	Permutedfs(nums, lenth, depth, used, path, &res) //深度递归
	return res
}
//深度递归遍历
func Permutedfs(nums []int, lenth int, depth int, used []bool, path []int, res *[][]int) {
	//terminator---终止递归的条件
	if depth == lenth {
		tmp := make([]int, lenth)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	//process
	//往中间结果集中加数据【此前未使用过的数据】
	for i := 0; i < lenth; i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			//接着遍历下一层
			Permutedfs(nums, lenth, depth+1, used, path, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
