/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	revInOrder(root, &sum)
	return root
}

func revInOrder(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	revInOrder(node.Right, sum)
	node.Val += *sum
	*sum = node.Val
	revInOrder(node.Left, sum)
}
