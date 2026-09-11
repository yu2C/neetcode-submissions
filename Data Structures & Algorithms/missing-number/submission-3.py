class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        xor = len(nums)

        for i, num in enumerate(nums):
            xor ^= i
            xor ^= num
        
        return xor
