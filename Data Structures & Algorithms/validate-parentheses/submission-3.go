func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	isLeft := map[rune]bool{'(': true, '{': true, '[': true}
	for _, c := range s {
		if isLeft[c] {
			stack = append(stack, c)
		} else if opener, ok := pairs[c]; ok {
			if len(stack) == 0 {
				return false
			}
			if opener == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}