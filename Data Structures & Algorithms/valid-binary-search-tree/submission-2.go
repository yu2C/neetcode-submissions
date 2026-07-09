/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return dfs(root, -1001, 1001)
}

func dfs(node *TreeNode, low int, high int) bool {
    if node == nil {
        return true
    }

    if !( low < node.Val && node.Val < high ) {
        return false
    }

    return dfs(node.Left, low, node.Val) && dfs(node.Right, node.Val, high)
}
