/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return convert(nums, 0, len(nums)-1)
}

func convert(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = convert(nums, left, mid-1)
	root.Right = convert(nums, mid+1, right)
	return root
}