class Solution:
    def numDecodings(self, s: str) -> int:
        memo: dict[int, int] = {}
        def dfs(i: int) -> int:
            if i == len(s):
                return 1
            
            if s[i] == '0':
                return 0

            if i in memo:
                return memo[i]
            
            res = dfs(i + 1)

            if i + 1 < len(s):
                val = int(s[i:i+2])

                if 10 <= val <= 26:
                    res += dfs(i+2)
            
            memo[i] = res
            return res
        
        return dfs(0)