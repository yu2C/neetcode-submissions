func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    res := make([]int, 0, n-k+1)

    // monotonic queue
    deque := make([]int, 0, k)

    for r := 0; r < n; r++ {
        if len(deque) > 0 && deque[0] <= r-k {
            deque = deque[1:]
        }

        for len(deque) > 0 &&
            nums[ deque[ len(deque) - 1 ] ] <= nums[r] {
                deque = deque[:len(deque)-1]
        }   

        deque = append(deque, r)

        if r >= k - 1 {
            res = append(res, nums[deque[0]])
        }
    }

    return res
}
