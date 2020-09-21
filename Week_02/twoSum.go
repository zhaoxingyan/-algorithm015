package Week_02
/*
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/
/*思路：
1、暴力求解，循环两遍数组(算法的复杂度为n*n)
2、hash表优化
*/
func twoSum(nums []int, target int) []int  {
	res := []int{}
	len := len(nums)

	if len <=1 {
		return res
	}

	for i := 0; i < len-1; i++ {
		for j := i+1; j < len; j++ {
			if nums[i] + nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				break
			}
		}
	}
	return res
}
/*
借助hash表实现，数组中的数为map的下标
遍历数组的数，如果存在2个数之和为target，其中一个数为
a在map中则target-a一定也在map中
*/
func twoSum2(nums []int, target int) []int{
	res := []int{}
	if len(nums) <= 1 {
		return res
	}
	midMap := make(map[int]int)
	for k,v := range nums {
		if val, ok := midMap[target-v]; ok {
			res = append(res, val)
			res = append(res, k)
			break
		} else {
			midMap[v] = k
		}
	}
	return res
}
