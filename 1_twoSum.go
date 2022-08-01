package leetcode

func twoSum(nums []int, target int) []int {
	n := make(map[int]int)
	for i, v := range nums {
		idx, ok := n[target-v]
		if ok {
			return []int{i, idx}
		}
		n[v] = i
	}
	return nil
}
