/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var MinVal, Prev int

func minDiffInBST(root *TreeNode) int {
	MinVal = math.MaxInt
	Prev = -1

	inOrder(root)
	return MinVal
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	if Prev != -1 {
		MinVal = min(MinVal, root.Val-Prev)
	}

	Prev = root.Val

	inOrder(root.Right)
}