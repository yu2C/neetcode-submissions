func lengthOfLIS(nums []int) int {
    // dp[i] = max(dp[i], dp[j]+1): longest increasing subsequence ending at nums[i]
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }
    ans := 1
    for i := range nums {
        for j := range i {
            if nums[j] < nums[i] {
                dp[i] = max(
                    dp[i],
                    dp[j]+1,
                )
            }
        }
        ans = max(ans, dp[i])
    }
    return ans
}
