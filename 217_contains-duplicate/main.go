package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, isSeen := set[v]; isSeen {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 10}))
}
