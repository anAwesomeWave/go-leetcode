package main

func productExceptSelf(nums []int) []int {
	totalProduct := 1
	totalProductWithoutZeroes := 1
	cntZeroes := 0

	for _, num := range nums {
		totalProduct *= num
		if num != 0 {
			totalProductWithoutZeroes *= num
		} else {
			cntZeroes += 1
		}
	}
	ans := make([]int, len(nums))

	for i, v := range nums {
		if v != 0 {
			ans[i] = totalProduct / v
		} else if cntZeroes < 2 {
			ans[i] = totalProductWithoutZeroes
		}
	}
	return ans
}
