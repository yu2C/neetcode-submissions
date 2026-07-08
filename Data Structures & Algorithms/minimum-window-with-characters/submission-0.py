class Solution:
    def minWindow(self, s: str, t: str) -> str:
        res = ""
        res_len = float("inf")

        l = 0

        need = Counter(t)
        need_count = len(need)
        have_count = 0
        win = defaultdict(int)
        res = [-1, -1]

        for r in range(len(s)):
            win[s[r]] += 1

            if s[r] in need and win[s[r]] == need[s[r]]:
                have_count += 1

            while have_count == need_count:
                if r-l+1 < res_len:
                    res = [l, r]
                    res_len = r-l+1
                
                win[s[l]] -= 1

                if s[l] in need and win[s[l]] < need[s[l]]:
                    have_count -= 1

                l += 1
        
        
        return "" if res_len == float("inf") else s[res[0]:res[1]+1]