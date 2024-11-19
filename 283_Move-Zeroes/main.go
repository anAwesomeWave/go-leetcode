package main

func moveZeroes(nums []int) {
	cnt := 0
	for i := range nums {
		if nums[i] == 0 {
			cnt += 1
			continue
		}
		nums[i-cnt] = nums[i]
	}
	for i := len(nums) - cnt; i < len(nums); i++ {
		nums[i] = 0
	}
}
