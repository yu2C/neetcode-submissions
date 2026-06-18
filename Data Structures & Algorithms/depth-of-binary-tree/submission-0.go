/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lMax := maxDepth(root.Left)
    rMax := maxDepth(root.Right)
    return max(lMax, rMax) + 1
}
