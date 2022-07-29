/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
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

	for _, c := range root.Children {
		helper(c, out)
	}

	*out = append(*out, root.Val)

}