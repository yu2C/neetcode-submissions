func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start := 0
	maxLen := 1

	expand := func(left int, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1

			if length > maxLen {
				start = left
				maxLen = length
			}

			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		// odd-length palindrome
		expand(i, i)
		// even-length palindrome
		expand(i, i+1)
	}

	return s[start:start+maxLen]
}
