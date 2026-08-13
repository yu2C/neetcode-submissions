class Solution:
    def foreignDictionary(self, words: List[str]) -> str:
        graph: dict[str, set[str]] = {}
        indegree: dict[str, int] = {}
        queue: deque[str] = deque()

        # preprocessing
        for word in words:
            for ch in word:

                graph.setdefault(ch, set())
                indegree.setdefault(ch, 0)
        
        for i in range(len(words)-1):
            word1 = words[i]
            word2 = words[i+1]

            foundDiff = False

            for ch1, ch2 in zip(word1, word2):
                if ch1 == ch2:
                    continue
                
                foundDiff = True

                if ch2 not in graph[ch1]:
                    graph[ch1].add(ch2)
                    indegree[ch2] += 1

                break
            
            if not foundDiff and len(word1) > len(word2):
                return ""
        
        # node -> queue
        for ch, degree in indegree.items():
            if degree == 0:
                queue.append(ch)
        
        # topological sort
        res = []

        while queue:
            node = queue.popleft()

            res.append(node)

            for neighbor in graph[node]:
                indegree[neighbor] -= 1

                if indegree[neighbor] == 0:
                    queue.append(neighbor)
        
        if len(res) != len(indegree):
            return ""
        
        return "".join(res)