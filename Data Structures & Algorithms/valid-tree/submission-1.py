class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        if n <= 0:
            return False
        if len(edges) != n - 1:
            return False
        
        graph = [[] for _ in range(n)]

        for node1, node2 in edges:
            if node1 < 0 or node1 >= n or node2 < 0 or node2 >= n:
                return False
            
            graph[node1].append(node2)
            graph[node2].append(node1)
        
        visited = [0] * n

        def dfs(node: int):
            visited[node] = True

            for neighbor in graph[node]:
                if visited[neighbor]:
                    continue
                dfs(neighbor)

        dfs(0)

        for node in range(n):
            if not visited[node]:
                return False
        return True
