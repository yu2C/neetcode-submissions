class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        hash_set = set(nums)

        long = 0

        for n in hash_set:
            l = 1

            while n+l in hash_set:
                l += 1
            
            long = max(long, l)
        
        return long


        
