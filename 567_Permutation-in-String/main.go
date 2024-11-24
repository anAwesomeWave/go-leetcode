package main

func isEqual(m1 [26]int, m2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	m1 := [26]int{}
	for i := range s1 {
		m1[s1[i]-'a'] += 1
	}
	m2 := [26]int{}

	len2 := 0
	for i := range s2 {
		if m1[s2[i]-'a'] > 0 {
			if len2 >= len(s1) {
				m2[s2[i-len2]-'a'] -= 1
				len2 -= 1
			}
			len2 += 1
			m2[s2[i]-'a'] += 1
			if len2 == len(s1) && isEqual(m1, m2) {
				return true
			}
		} else {
			len2 = 0
			m2 = [26]int{}
		}
	}
	return false
}
