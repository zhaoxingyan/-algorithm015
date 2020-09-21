package Week_02
/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

 

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
 

提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
 */
/*
思路：堆排序的典型用法就是topK问题
1、遍历数组对元素出现的次数进行统计存放到hash表中
2、取出hash表中的值存放到数组中，针对数组中的元素构造容量为K的小顶堆
3、对数组中的其它元素做遍历与小顶堆的顶部元素做对比，小于顶部元素直接舍弃，超过小顶堆的数据
需要移除原小顶堆的数据将新的数据加入，最后的结果就是topK的频次数
4、遍历原hash表，将hash表中的k加入结果集
该算法的时间复杂度主要是在建堆以及对堆进行排序上O(nlogk)
*/
func topKFrequent(nums []int, k int) []int {
	var res []int
	mapS := make(map[int]int)
	for _, v := range nums {
		if val, ok := mapS[v]; ok {
			mapS[v] = val + 1
		} else {
			mapS[v] += 1
		}
	}
	var mid []int
	for _, v := range mapS {
		mid = append(mid, v)
	}
	//构建容量为K的小顶堆
	heap := NewHeap(k)
	for i := 0; i < k; i++ {
		heap.Insert(mid[i])
	}
	for i := k; i < len(mid); i++ {
		if mid[i] > heap.a[1] {
			heap.removeSmall()
			heap.Insert(mid[i])
		}
	}
	a := heap.a[1:]
	for _, v := range a {
		for sk, sv := range mapS {
			if sv == v && !isUsed(res, sk) {
				res = append(res, sk)
			}
		}
	}
	return res
}

func isUsed(a []int, k int) bool {
	if len(a) <= 0 {
		return false
	}
	for _, v := range a {
		if k == v {
			return true
		}
	}
	return false
}
