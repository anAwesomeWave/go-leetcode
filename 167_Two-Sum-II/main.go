package main

func bs(arr []int, target int) int {
	// pos of found target, -1 if not found

	l, r := 0, len(arr)-1
	if arr[r] == target {
		return r
	}
	for l < r-1 {
		m := l + (r-l)/2
		if arr[m] > target {
			r = m
		} else {
			l = m
		}
	}
	if arr[l] == target {
		return l
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	for ind, x := range numbers {
		ans := bs(numbers, target-x)
		if ans != -1 && ans != ind {
			return []int{ind + 1, ans + 1}
		}
	}
	return nil
}
