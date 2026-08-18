class Solution:
    def rob(self, nums: List[int]) -> int:
        prev1, prev2 = 0, 0

        for i in range(len(nums)):
            cur = max(prev1, prev2 + nums[i])
        
            prev2 = prev1
            prev1 = cur
    
        return prev1