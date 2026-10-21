func isAnagram(s string, t string) bool {
	numMap := make(map[rune]int)

	for _, char := range s {
		numMap[char]++
	}
	for _, char := range t {
		numMap[char]--
	}

	for _, v := range numMap {
		if v != 0 {
			return false
		}
	}
	return true
}
