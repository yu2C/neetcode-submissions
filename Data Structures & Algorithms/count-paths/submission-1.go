func uniquePaths(m int, n int) int {
	memo := make(map[[2]int]int)
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}

		if i == m-1 && j == n-1 {
			return 1
		}

		key := [2]int{i, j}

		if val, ok := memo[key]; ok {
			return val
		}

		memo[key] = dfs(i+1, j) + dfs(i, j+1)

		return memo[key]
	}

	return dfs(0, 0)
}
