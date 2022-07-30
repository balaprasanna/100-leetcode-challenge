/**
https://leetcode.com/problems/maximum-depth-of-binary-tree/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Max(x,y int) int {
    if (x > y) {
        return x
    } else {
        return y
    }
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } 
    return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}
