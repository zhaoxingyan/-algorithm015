package Week_01

//删除排序数组中的重复项
func RemoveDuplicates(nums []int) int{
	len := len(nums)
	count := 0
	for i := 1; i < len; i++ {
		if nums[i] != nums[count] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}

//旋转数组
func Rotate(nums []int, k int)  {
	//1、反转数组
	//2、反转前k%len个数
	//3、反转后n-k%k个数
	len := len(nums)
	step := k % len
	reverse(nums, 0, len-1)
	reverse(nums, 0, step-1)
	reverse(nums, step, len-1)
}
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
//两数之和
func woSum(nums []int, target int) []int {
	res := []int{}
	data := make(map[int]int)
	for i, v := range nums {
		if  val, ok := data[target - v]; ok {
			res = append(res, val)
			res = append(res, i)
			break
		} else {
			data[v] = i
		}
	}
	return res
}
//移动零
func MoveZeroes(nums []int) []int {
	j := 0
	len := len(nums)
	if len <= 1 {
		return nums
	}
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			nums[i],nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums
}
//加一
func PlusOne(digits []int) []int {
	len := len(digits)
	for i := len - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	digits[0] = 1
	return append(digits, 0)

}
//合并2个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int)  {
	p1 := m - 1
	p2 := n -1
	p := m + n - 1
	for p2 >= 0 {
		if p1 >= 0 && (nums1[p1] > nums2[p2]) {
			nums1[p] = nums1[p1]
			p--
			p1--
		} else {
			nums1[p] = nums2[p2]
			p--
			p2--
		}
	}
}


