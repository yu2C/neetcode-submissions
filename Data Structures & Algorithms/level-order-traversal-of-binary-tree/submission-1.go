/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)
        lvl := make([]int, 0)

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            lvl = append(lvl, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, lvl)
    }
    return res
}
