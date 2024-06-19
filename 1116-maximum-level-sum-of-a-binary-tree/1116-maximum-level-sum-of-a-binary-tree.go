/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return -1
	}
	maxLevel := 1
	level := 1
	maxSum := root.Val
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0

		for i := 0; i < size; i++ {
			node := queue[i]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			maxLevel = level

		}
		queue = queue[size:]
		level++
	}
	return maxLevel
}