class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        rows = len(matrix)
        cols = len(matrix[0])

        res = []

        top, bottom = 0, rows - 1
        left, right = 0, cols - 1

        while top <= bottom and left <= right:
            # top: left -> right
            for col in range(left, right+1):
                res.append(matrix[top][col])
            top += 1

            # right: top -> bottom
            for row in range(top, bottom+1):
                res.append(matrix[row][right])
            right -= 1

            if top <= bottom:
                # bottom: right -> left
                for col in range(right, left-1, -1):
                    res.append(matrix[bottom][col])
                bottom -= 1
            if left <= right:
                # left: bottom -> top
                for row in range(bottom, top-1, -1):
                    res.append(matrix[row][left])
                left += 1
        return res