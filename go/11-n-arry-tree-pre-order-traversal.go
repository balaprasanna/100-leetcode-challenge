/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	out := []int{}
	if root == nil {
		return out
	}
	helper(root, &out)
	return out
}

func helper(root *Node, out *[]int) {
	if root == nil {
		return
	}

	*out = append(*out, root.Val)

	for _, c := range root.Children {
		helper(c, out)
	}

}