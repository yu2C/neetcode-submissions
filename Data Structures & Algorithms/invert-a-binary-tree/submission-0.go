/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return root
}