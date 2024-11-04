package main

func isAnagram(s string, t string) bool {
	// time O(n + m)
	// space O(n)
	// for alphabetical chars only O(1) -> we can only store slice with 26 letters
	charsMap := map[rune]int{}

	for _, char := range s {
		charsMap[char] += 1
	}
	for _, char := range t {
		charsMap[char] -= 1
	}
	for _, v := range charsMap {
		if v != 0 {
			return false
		}
	}
	return true
}
