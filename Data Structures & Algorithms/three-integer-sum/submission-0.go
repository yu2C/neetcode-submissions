func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] { continue }
        s, e := i+1, len(nums)-1
        for s < e {
            sum := nums[i] + nums[s] + nums[e]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[s], nums[e]})
                for s < e && nums[s] == nums[s+1] { s++ }
                for s < e && nums[e] == nums[e-1] { e-- }
                s++; e--
            } else if sum < 0 {
                s++
            } else {
                e--
            }
        }
    }
    return res
}