class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        rows = len(heights)
        cols = len(heights[0])

        pac: list[list[bool]] = [[False] * cols for _ in range(rows)]
        atl: list[list[bool]] = [[False] * cols for _ in range(rows)]

        directions: list[tuple[int, int]] = [
            (-1, 0),
            (1, 0),
            (0, -1),
            (0, 1),
        ]

        def dfs(row: int, col: int, visited: list[list[bool]]):
            visited[row][col] = True

            for delta_row, delta_col in directions:
                next_row = row + delta_row
                next_col = col + delta_col

                if next_row < 0 or next_row >= rows or next_col < 0 or next_col >= cols:
                    continue
                
                if visited[next_row][next_col]:
                    continue

                if heights[next_row][next_col] < heights[row][col]:
                    continue
                
                dfs(next_row, next_col, visited)
        
        for col in range(cols):
            dfs(0, col, pac)
            dfs(rows-1, col, atl)
        
        for row in range(rows):
            dfs(row, 0, pac)
            dfs(row, cols-1, atl)

        res = []

        for row in range(rows):
            for col in range(cols):
                if pac[row][col] and atl[row][col]:
                    res.append((row, col))
        
        return res