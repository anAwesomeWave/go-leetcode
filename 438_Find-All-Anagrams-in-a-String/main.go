package main

func isEqual(m1, m2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	m1 := [26]int{}
	for i := range p {
		m1[p[i]-'a'] += 1
	}
	m2 := [26]int{}
	len2 := 0
	ans := []int{}
	for i := range s {
		if m1[s[i]-'a'] > 0 {
			if len2 >= len(p) {
				m2[s[i-len2]-'a'] -= 1
				len2 -= 1
			}
			len2 += 1
			m2[s[i]-'a'] += 1
			if len2 == len(p) && isEqual(m1, m2) {
				ans = append(ans, i-len2+1)
			}

		} else {
			m2 = [26]int{}
			len2 = 0
		}
	}
	return ans
}
