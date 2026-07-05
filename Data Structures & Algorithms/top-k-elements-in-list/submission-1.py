class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = Counter(nums)
        
        pairs = []

        for key, value in count.items():
            pairs.append((key, value))
        
        pairs.sort(key=lambda x: x[1], reverse=True)

        return [pairs[i][0] for i in range(k)]