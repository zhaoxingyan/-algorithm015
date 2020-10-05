package Week_03

import "sort"

func permuteUnique(nums []int) [][]int {
	//排序
	lenth := len(nums)
	if lenth < 0 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	var path []int

	used := make([]bool, lenth)
	permuteUniqueDfs(nums, lenth, used, 0, path, &res)
	return res
}

func permuteUniqueDfs(nums []int, lenth int, used []bool, depth int, path []int, res *[][]int) {
	//terminator
	if depth == lenth {
		tmp := make([]int, lenth)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < lenth; i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i - 1] {
			continue
		}
		path = append(path, nums[i]) //加入路径中
		used[i] = true
		permuteUniqueDfs(nums, lenth, used, depth+1, path, res)
		used[i] = false
		path = path[:len(path)-1]   //回溯
	}
}
