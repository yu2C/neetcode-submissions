class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
         nums = sorted(nums)
         j = 0
         for i in range(1, len(nums)):
            if nums[j] == nums[i]:
                return True
            j += 1
         return False
