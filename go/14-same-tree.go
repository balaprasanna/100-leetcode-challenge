/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if (p == nil && q == nil) {
        return true
    }
    return helper(p,q)
}

func helper(p *TreeNode, q *TreeNode) bool {
    if (p == nil && q == nil) {
        return true
    } else if ((p == nil && q != nil) || (p != nil && q == nil)) {
        return false
    } else if (p.Val != q.Val) {
        return false
    } else {
        return helper(p.Left, q.Left) && helper(p.Right, q.Right)
    }
    
}
