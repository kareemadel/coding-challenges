/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return maxDepthHelper(root, 0)
}

func maxDepthHelper(n *TreeNode, depth int) int {
	if n == nil {
		return depth
	}
	return max(maxDepthHelper(n.Left, depth+1), maxDepthHelper(n.Right, depth+1))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}