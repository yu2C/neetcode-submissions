func countBits(n int) []int {
    dp := make([]int, n+1)
    dp[0] = 0

    for i := 1; i <= n; i++ {
        dp[i] = dp[i & (i-1)] + 1
    }
    return dp
}
