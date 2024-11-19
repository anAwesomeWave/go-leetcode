package main

/*
суть: мы на каждом шаге накопили сумму на отрезке [0, curInd], но этого может быть много/мало
для того, чтобы набрать target=k. Посмотрим, есть ли среди встреченных нужная избыточная, недостаточная сумма.
тогда мы сможем отрубить такое начало (столько элементов)
*/
func subarraySum(nums []int, k int) int {
	prefSum := make(map[int]int) // sum -> count
	curSum := 0
	ans := 0
	prefSum[0] = 1
	for _, v := range nums {
		curSum += v
		ans += prefSum[curSum-k]
		prefSum[curSum] += 1

	}
	return ans
}
