package Week_02

/*
242  有效的字母异位词
https://leetcode-cn.com/problems/valid-anagram/
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true

题解思路：字母异位词（两个单词中字母出现的数量是一致的）
两种解题思路：1、 对单词排序，排序后比较两个是否一致
            2、hashmap记录每个单词中字母出现的次数，比较两个map
*/
func isAnagram(s string, t string) bool {
	//异位词的长度必然相等
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[uint8]int)
	mapT := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if count1,ok := mapS[s[i]];ok{
			mapS[s[i]] = count1 + 1
		} else {
			mapS[s[i]] = 1
		}
		if count2,ok := mapT[t[i]];ok{
			mapT[t[i]] = count2 + 1
		} else {
			mapT[t[i]] = 1
		}
	}
	for k,v := range mapS {
		if count, ok := mapT[k]; ok  && count == v{
			continue
		}else {
			return false
		}
	}
	return true
}
//更高效解法
/*
一个长度为26的数组来保存位数不相等的字母，
s中有某字母该字母数加一
t中有该字母该字母数减一
最后遍历改数组，数组中有数据不为零则不是异位字母组
*/
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mid := make([]int, 26)//保存字母数量的数组
	for i := 0; i < len(s); i++ {
		mid[s[i]-'a']++
		mid[t[i]-'a']--
	}
	for _,v := range mid {
		if v != 0 {
			return false
		}
	}
	return true
}