package main

func groupAnagrams(strs []string) [][]string {
	// Time O(n * m), n - len(strs), m - 100 (len of word)
	// Space O(n) - n len(strs)

	anagrams := map[[26]int][]string{}
	for _, curStr := range strs {
		//r := []rune(curStr)                 // O(n = len(curStr))
		//sort.Slice(r, func(i, j int) bool { // O(nlogn)
		//	return r[i] < r[j]
		//})
		//fmt.Println(string(r)) // O(n) -- too expensive approach

		// slice cannot be key in map, but arr26 can
		lettersCnt := [26]int{}
		for _, lt := range curStr {
			lettersCnt[lt-'a'] += 1
		}
		anagrams[lettersCnt] = append(anagrams[lettersCnt], curStr)
	}
	var ans [][]string
	for _, anagram := range anagrams {
		ans = append(ans, anagram)
	}
	return ans
}
