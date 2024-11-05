package main

func isAlphaNumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func toLower(c byte) byte {
	var gap byte = 'a' - 'A' // should be 32
	if c >= 'A' && c <= 'Z' {
		return c + gap
	}
	return c
}

func isPalindrome(s string) bool {
	cur := 0
	rcur := len(s) - 1
	for cur < rcur {
		if !isAlphaNumeric(s[cur]) {
			cur += 1
			continue
		}
		if !isAlphaNumeric(s[rcur]) {
			rcur -= 1
			continue
		}
		if toLower(s[cur]) != toLower(s[rcur]) {
			return false
		}
		cur += 1
		rcur -= 1
	}
	return true
}
