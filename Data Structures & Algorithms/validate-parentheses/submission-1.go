import "slices"

func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	left := []rune{'(', '{', '['}
	for _, c := range s {
		if slices.Contains(left, c) {
			stack = append(stack, c)
		}
        if len(stack) == 0 { return false }
		if opener, ok := pairs[c]; ok {
			if opener == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
