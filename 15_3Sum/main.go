package main

import "sort"

func threeSum(nums []int) [][]int {
	// алгос такой, берем каждое число, считаем, что его точно берем.
	// для него рассматриваем все следующие числа,
	// для каждой пары смотрим в мапу с предыдущими числами
	// все 3 числа различные. 1 число посередине, 2 справа, 3 слева

	// можно посортить и использовать два указателя для каждого числа, но ассимптотика будет также O(n^2)
	ans := [][]int{}
	seenNums := map[int]struct{}{}        // mark all prev senn nums
	usedTriplets := map[[3]int]struct{}{} // used triplets
	for ind, val := range nums {
		for _, secVal := range nums[ind+1:] {
			if _, ok := seenNums[-(val + secVal)]; ok {

				triplet := []int{val, secVal, -(val + secVal)}
				newKey := [3]int{}
				sort.Slice(
					triplet,
					func(i, j int) bool { return triplet[i] < triplet[j] },
				)
				copy(newKey[:], triplet)
				if _, ok := usedTriplets[newKey]; !ok {
					ans = append(ans, triplet)
					usedTriplets[newKey] = struct{}{}
				}
			}
		}
		seenNums[val] = struct{}{}
	}
	return ans
}
