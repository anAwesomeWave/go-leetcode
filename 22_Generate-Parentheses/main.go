package main

func recGen(curStr []byte, totalOpened int, curPair int, n int, ans *[]string) {
	if curPair == n && totalOpened == 0 {
		(*ans) = append(*ans, string(curStr))
		return
	}
	if totalOpened > 0 {
		clCurStr := append(curStr, ')')
		recGen(clCurStr, totalOpened-1, curPair, n, ans)
	}
	if curPair < n {
		curStr = append(curStr, '(')
		recGen(curStr, totalOpened+1, curPair+1, n, ans)
	}
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0, n*n)
	str := make([]byte, 0, n*2)
	recGen(str, 0, 0, n, &ans)
	return ans
}
