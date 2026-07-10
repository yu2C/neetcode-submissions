# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        v = []
        def bfs(node):
            if not node:
                return
            
            bfs(node.left)
            v.append(node.val)
            bfs(node.right)
        bfs(root)
        return v[k-1]