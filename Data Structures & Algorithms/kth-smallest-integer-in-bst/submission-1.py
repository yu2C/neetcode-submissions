# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        ans = None

        def m(node):
            nonlocal k, ans

            if not node or ans is not None:
                return

            m(node.left)

            k -= 1
            if k == 0:
                ans = node.val
                return
            
            m(node.right)

        m(root)
        return ans