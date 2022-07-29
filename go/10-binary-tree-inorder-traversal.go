/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	output := []int{}
	if root == nil {
		return output
	} else {
		helper(root, &output)
	}
	return output
}

func helper(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, output)
	*output = append(*output, root.Val)
	helper(root.Right, output)
}