func coinChange(coins []int, amount int) int {
    // dp[i] the least number of coins can fits amount
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}

			dp[i] = min(
				dp[i],
				dp[i-coin]+1,
			)
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}
