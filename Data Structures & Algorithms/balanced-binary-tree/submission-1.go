/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) (bool, int)

    dfs = func(node *TreeNode) (bool, int) {
        if node == nil {
            return true, 0
        }

        lB, lH := dfs(node.Left)
        rB, rH := dfs(node.Right)

        b := lB && rB && abs(lH - rH) <= 1
        return b, max(lH, rH)+1
    }
    b, _ := dfs(root)
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}