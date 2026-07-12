class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        res: List[List[int]] = []
        cur: List[int] = []

        def dfs(idx: int, rem: int):
            if rem == 0:
                res.append(cur.copy())
                return
            
            if rem < 0 or idx == len(nums):
                return
            
            cur.append(nums[idx])
            dfs(idx, rem-nums[idx])
            cur.pop()

            dfs(idx+1, rem)

        dfs(0, target)

        return res