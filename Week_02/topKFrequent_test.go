package Week_02

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	num := []int{1,1,1,1,5,5,5,6,6,1,3,3,3}
	k := 1
	fmt.Println(topKFrequent(num, k))
}
