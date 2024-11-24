package main

func isOpen(bracket byte) bool {
	for _, br := range []byte{'(', '{', '['} {
		if bracket == br {
			return true
		}
	}
	return false
}

func matchingBracket(bracket byte) byte {
	switch bracket {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	}
	return ' '
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if isOpen(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != matchingBracket(s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
