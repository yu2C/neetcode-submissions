func maxProduct(nums []int) int {
    n := len(nums)

    maxDp := make([]int, n)
    minDp := make([]int, n)

    maxDp[0] = nums[0]
    minDp[0] = nums[0]

    res := nums[0]

    for i := 1; i < n; i++ {
        x := nums[i]

        maxDp[i] = max(
            x,
            maxDp[i-1] * x,
            minDp[i-1] * x,
        )
        minDp[i] = min(
            x,
            maxDp[i-1] * x,
            minDp[i-1] * x,
        )
        res = max(res, maxDp[i])
    }
    return res
}
