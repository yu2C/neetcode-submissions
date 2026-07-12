func combinationSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    cur := make([]int, 0)

    var dfs func(idx int, rem int)

    dfs = func(idx int, rem int) {
        if rem == 0 {
            com := append([]int(nil), cur...)
            res = append(res, com)
            return
        }

        if rem < 0 || idx == len(nums) {
            return
        }

        cur = append(cur, nums[idx])
        dfs(idx, rem - nums[idx])
        cur = cur[:len(cur)-1]

        dfs(idx+1, rem)
    }        
    dfs(0, target)
    return res
}
