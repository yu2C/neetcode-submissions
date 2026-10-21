func isPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1

	for l < r {
		for l < r && !isAl(s[l]) {
			l++
		}
		for l < r && !isAl(s[r]) {
			r--
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}
	return true
}

func isAl(c byte) bool {
	return (c >= 'a' && c <= 'z' ) ||
		( c >= 'A' && c <= 'Z' ) ||
		( c >= '0' && c <= '9' )
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ( 'a' - 'A' )
	}
	return c
}