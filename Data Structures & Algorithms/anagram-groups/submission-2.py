class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = []
        m = defaultdict(list)

        for s in strs:
            count = [0] * 26
            for char in s:
                count[int(ord(char) - ord('a'))] += 1
            m[tuple(count)].append(s)
        for v in m.values():
            res.append(v)
        return res