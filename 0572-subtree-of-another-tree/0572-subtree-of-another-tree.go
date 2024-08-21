/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isIdentical(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isIdentical(root1.Left, root2.Left) && isIdentical(root1.Right, root2.Right)
}