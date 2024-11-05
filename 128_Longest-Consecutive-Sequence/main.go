package main

func dfs(cur int, gr map[int]bool) int {
	gr[cur] = true
	ans := 1
	if v, ok := gr[cur-1]; ok && !v {
		ans += dfs(cur-1, gr)
	}

	if v, ok := gr[cur+1]; ok && !v {
		ans += dfs(cur+1, gr)
	}

	return ans
}

func longestConsecutive(nums []int) int {
	gr := make(map[int]bool, len(nums))
	for _, v := range nums {
		gr[v] = false
	}
	ans := 0
	for k, v := range gr {
		if !v {
			curAns := dfs(k, gr)
			if curAns > ans {
				ans = curAns
			}
		}
	}
	return ans
}
