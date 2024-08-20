/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	i := 0
	return inOrder(root, k, &i)
}

func inOrder(root *TreeNode, k int, i *int) int {
	if root == nil {
		return -1
	}

	val := inOrder(root.Left, k, i)
	if val != -1 {
		return val
	}

	*i++
	if *i == k {
		return root.Val
	}
	return inOrder(root.Right, k, i)
}