package Week_02

import (
	"fmt"
	"testing"
)

func TestIsUgly(t *testing.T)  {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(-5))
	fmt.Println(isUgly(14))
}
