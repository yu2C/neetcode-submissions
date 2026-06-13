class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        for i in range(len(nums) - 2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            s, e = i + 1, len(nums)-1
            while s < e:
                summ = nums[i] + nums[s] + nums[e]
                if summ == 0:
                    res.append([nums[i], nums[s], nums[e]])
                    while s < e and nums[s] == nums[s+1]:
                        s += 1
                    while s < e and nums[e] == nums[e-1]:
                        e -= 1
                    s += 1
                    e -= 1
                elif summ < 0:
                    s += 1
                else:
                    e -= 1
        return res