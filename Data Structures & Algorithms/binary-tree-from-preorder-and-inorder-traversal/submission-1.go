/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    m := make(map[int]int)
    for i := range inorder {
        m[inorder[i]] = i
    }

    var helper func(preL, preR, inL, inR int) *TreeNode
    helper = func(preL, preR, inL, inR int) *TreeNode {
        if preL > preR {
            return nil
        }

        rootV := preorder[preL]
        node := &TreeNode{ Val:rootV }

        rootI := m[rootV]
        lSize := rootI - inL

        node.Left = helper(
            preL + 1,
            preL + lSize,
            inL,
            rootI - 1,
        )

        node.Right = helper(
            preL + lSize + 1,
            preR,
            rootI + 1,
            inR,
        )
        return node
    }
    return helper(0, len(preorder)-1, 0, len(inorder)-1)
}
