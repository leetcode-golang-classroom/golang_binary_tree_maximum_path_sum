package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := root.Val
	MaxSum(root, &result)
	return result
}

func MaxSum(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	leftMax := MaxSum(root.Left, result)
	rightMax := MaxSum(root.Right, result)
	// choose or not choose
	leftMax = Max(leftMax, 0)
	rightMax = Max(rightMax, 0)
	// split
	*result = Max(*result, root.Val+leftMax+rightMax)
	return root.Val + Max(leftMax, rightMax)
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
