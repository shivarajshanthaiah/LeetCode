/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{}
	dfs(root, &nodes)
	return balance(nodes)
}

func dfs(node *TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		dfs(node.Left, nodes)
	}
	*nodes = append(*nodes, node)
	if node.Right != nil {
		dfs(node.Right, nodes)
	}
}

func balance(nodes []*TreeNode) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	if len(nodes) == 1 {
		nodes[0].Left, nodes[0].Right = nil, nil
		return nodes[0]
	}
	newRoot := nodes[len(nodes)/2]
	newRoot.Left = balance(nodes[:len(nodes)/2])
	newRoot.Right = balance(nodes[len(nodes)/2+1:])
	return newRoot
}