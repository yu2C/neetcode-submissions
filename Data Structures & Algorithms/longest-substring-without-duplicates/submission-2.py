class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        l = 0
        res = 0

        m = {}

        for r in range(len(s)):
            while s[r] in m:
                del m[s[l]]
                l += 1
            m[s[r]] = True

            res = max(res, r-l+1)
        return res