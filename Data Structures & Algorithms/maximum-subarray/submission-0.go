func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// dp[i] = max(dp[i-1] + nums[i], nums[i])
	curSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums);i++ {
		curSum = max(
			curSum + nums[i],
			nums[i],
		)

		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
