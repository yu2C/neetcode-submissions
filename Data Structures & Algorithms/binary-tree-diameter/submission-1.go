/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    ans := 0
    var h func(node *TreeNode) int 
    h = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        l := h(node.Left)
        r := h(node.Right)
        ans = max(ans, l + r)

        return max(l, r) + 1
    }
    h(root)
    return ans
}