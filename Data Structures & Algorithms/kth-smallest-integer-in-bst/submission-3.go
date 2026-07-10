/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    ans := -1
    m(root, &ans, &k)
    return ans
}

func m(node *TreeNode, ans *int, k *int) {
    if node == nil || *ans != -1 {
        return
    }

    m(node.Left, ans, k)

    *k--
    if *k == 0 {
        *ans = node.Val
        return
    }
    m(node.Right, ans, k)
}