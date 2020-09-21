package Week_02

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	k := "agramna"
	fmt.Println(isAnagram(s, k))
	fmt.Println(isAnagram("rat", "cat"))
}

func TestIsAnagram2(t *testing.T) {
	fmt.Println(isAnagram2("anagram", "agramna"))
	fmt.Println(isAnagram2("rat", "cat"))
}