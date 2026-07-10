/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    v := []int{}
    m(root, &v)
    return v[k-1]
}

func m(node *TreeNode, v *[]int) {
    if node == nil {
        return
    }

    m(node.Left, v)
    *v = append(*v, node.Val)
    m(node.Right, v)
}