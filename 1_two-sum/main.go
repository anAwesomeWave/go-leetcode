package main

func twoSum(nums []int, target int) []int {
	// Time: O(n)
	// Space: O(n)
	valToInd := map[int]int{} // value to index in nums

	for ind, val := range nums {
		if secondInd, ok := valToInd[target-val]; ok {
			// found pair
			return []int{ind, secondInd}
		}
		valToInd[val] = ind
	}
	return []int{-1, -1}
}
