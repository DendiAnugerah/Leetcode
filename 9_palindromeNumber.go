package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {
	C := strconv.Itoa(x)
	l := len(C)
	if x < 0 {
		return false
	}
	for i := 0; i < len(C)/2; i++ {
		if C[i] == C[l-1] {
			l--
			continue
		}
		return false
	}
	return true
}
