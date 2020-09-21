package Week_02

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	num := []int{2, 7, 11, 15}
	target := 18
	res := twoSum(num, target)
	fmt.Println(res)
}

func TestTwoSum2(t *testing.T) {
	num := []int{2, 7, 11, 15}
	target := 18
	res := twoSum2(num, target)
	fmt.Println(res)
}
