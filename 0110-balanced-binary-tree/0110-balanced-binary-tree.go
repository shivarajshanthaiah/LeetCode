/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.Left)
	right := height(root.Right)

	if abs(left-right) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := height(root.Left)
	right := height(root.Right)

	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}