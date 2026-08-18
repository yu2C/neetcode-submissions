func rob(nums []int) int {
    n := len(nums)

    if n == 1 {
        return nums[0]
    }

    robRange := func(start int, end int) int {
        prev1, prev2 := 0, 0

        for i := start; i <= end; i++ {
            cur := max(prev1, prev2+nums[i])
            prev2 = prev1
            prev1 = cur
        }
        return prev1
    }

    woLast := robRange(0, n-2)
    woFirst := robRange(1, n-1)

    return max(woLast, woFirst)
}
