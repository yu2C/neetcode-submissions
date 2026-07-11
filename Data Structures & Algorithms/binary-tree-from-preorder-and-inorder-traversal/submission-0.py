# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        m = { v: i for i, v in enumerate(inorder) }
        def helper(pre_l, pre_r, in_l, in_r):
            if pre_l > pre_r:
                return None
            root_val = preorder[pre_l]
            node = TreeNode(root_val)
            
            root_i = m[root_val]
            l_size = root_i - in_l

            node.left = helper(
                pre_l + 1,
                pre_l + l_size,
                in_l,
                root_i - 1
                )
            node.right = helper(
                pre_l + l_size + 1,
                pre_r,
                root_i + 1,
                in_r
            )

            return node
        return helper(0, len(preorder)-1, 0, len(inorder)-1)