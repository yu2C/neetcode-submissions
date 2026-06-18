# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        ans = 0

        def height(node: [TreeNode]) -> int:
            nonlocal ans

            if not node:
                return 0

            l_h = height(node.left)
            r_h = height(node.right)

            ans = max(ans, l_h + r_h)

            return max(l_h, r_h) + 1
        height(root)
        return ans