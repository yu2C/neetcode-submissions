class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        graph = [[] for _ in range(n)]

        for edge in edges:
            graph[edge[0]].append(edge[1])
            graph[edge[1]].append(edge[0])
        
        visited = [False] * n

        def dfs(node: int):
            visited[node] = True

            for neighbor in graph[node]:
                if visited[neighbor]:
                    continue
                
                dfs(neighbor)
        
        count = 0

        for node in range(n):
            if visited[node]:
                continue
            count += 1
            dfs(node)
        
        return count