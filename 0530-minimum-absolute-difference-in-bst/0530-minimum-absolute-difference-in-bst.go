/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var Res, PrevVal int

func getMinimumDifference(root *TreeNode) int {
	Res = math.MaxInt
	PrevVal = -1

	inOrder(root)
	return Res
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	if PrevVal != -1 {
		Res = min(Res, root.Val-PrevVal)
	}
    
	PrevVal = root.Val

	inOrder(root.Right)
}