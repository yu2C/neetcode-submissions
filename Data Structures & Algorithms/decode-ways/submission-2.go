func numDecodings(s string) int {
    memo := make(map[int]int)

    var dfs func(i int) int
    dfs = func(i int) int {
        if i == len(s) {
            return 1
        }

        if s[i] == '0' {
            return 0
        }

        if val, exists := memo[i]; exists {
            return val
        }

        res := dfs(i + 1)

        if i + 1 < len(s) {
            val, err := strconv.Atoi(s[i:i+2])
            if err != nil {
                return 0
            }

            if 10 <= val && val <= 26 {
                res += dfs(i+2)
            }
        }
        memo[i] = res
        return res
    }
    return dfs(0)
}
