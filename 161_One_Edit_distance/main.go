package main

func abs[T int | int32 | int64](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}
func IsOneEditDistance(s string, t string) bool {
	distLen := abs(len(s) - len(t))
	if distLen > 1 {
		return false
	}
	if distLen == 0 {
		mismatched := 0
		for i, _ := range s {
			if s[i] != t[i] {
				mismatched += 1
			}
		}
		return mismatched != 0 && mismatched < 2
	} else {
		// len(s) == len(t) + 1 or otherwise
		if len(s) > len(t) {
			// deletion
			deleted := 0
			for i, _ := range s {
				if i-deleted >= len(t) {
					continue
				}
				if s[i] != t[i-deleted] {
					deleted += 1
				}
			}
			return deleted < 2
		} else {
			added := 0
			for i, _ := range s {
				if s[i-added] != t[i] {
					added += 1
				}
			}
			return added < 2
		}
	}
	// write your code here
}
