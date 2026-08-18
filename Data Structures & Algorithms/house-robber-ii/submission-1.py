class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 0:
            return 0
        if n == 1:
            return nums[0]
        return max(
            self.robRange(nums[:n-1]),
            self.robRange(nums[1:]),
        )
    
    def robRange(self, nums: List[int]) -> int:
        prev1, prev2 = 0, 0

        for num in nums:
            cur = max(prev1, prev2+num)
            prev2 = prev1
            prev1 = cur
        
        return prev1