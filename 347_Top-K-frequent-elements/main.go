package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Time O(n)
	// Space O(n + k)

	freqs := map[int]int{}
	rFreqs := map[int][]int{} // reversed. freq-> nums
	for _, num := range nums {
		freqs[num] += 1
	}
	for num, freq := range freqs {
		rFreqs[freq] = append(rFreqs[freq], num)
	}
	ans := make([]int, 0, k)

	for i := len(nums); i > 0; i-- {
		for _, j := range rFreqs[i] {
			ans = append(ans, j)
			if len(ans) == k {
				return ans
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(topKFrequent([]int{1, 2, 3}, 3))
}
