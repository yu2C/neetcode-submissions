class Solution:
	def uniquePaths(self, m: int, n: int) -> int:
		memo = [[0] *n for _ in range(m)]

		def dfs(i: int, j: int) -> int:
			if i >= m or j >= n:
				return 0
			if i == m-1 and j == n-1:
				return 1
			
			if memo[i][j] != 0:
				return memo[i][j]

			memo[i][j] = dfs(i, j+1) + dfs(i+1, j)

			return memo[i][j]
		
		return dfs(0, 0)