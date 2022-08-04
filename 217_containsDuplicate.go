package leetcode

func containsDuplicate(nums []int) bool {
	counter := make(map[int]bool)
	for _, val := range nums {
		if counter[val] {
			return true
		} else {
			counter[val] = true
		}
	}
	return false
}
