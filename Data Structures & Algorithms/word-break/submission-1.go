func wordBreak(s string, wordDict []string) bool {
    n := len(s)

	dp := make([]bool, n+1)
	dp[0] = true

	words := make(map[string]bool, len(wordDict))

	for _, word := range wordDict {
		words[word] = true
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			_, exists := words[s[j:i]]
			if dp[j] && exists {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
