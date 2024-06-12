/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	node := searchBST(root.Right, val)
	if node != nil {
		return node
	}

	node = searchBST(root.Left, val)
	if node != nil {
		return node
	}

	return nil
}