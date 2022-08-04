package leetcode

func moveZeroes(nums []int) {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ans], nums[i] = nums[i], nums[ans]
			ans++
		}
	}
}
