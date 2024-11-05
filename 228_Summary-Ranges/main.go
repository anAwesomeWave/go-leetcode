package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ans := []string{}
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if nums[i]-1 != nums[i-1] {
			if start == nums[i-1] {
				ans = append(ans, fmt.Sprintf("%d", start))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = v
		}
	}
	if nums[len(nums)-1] == start {
		ans = append(ans, fmt.Sprintf("%d", start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	}
	return ans
}
