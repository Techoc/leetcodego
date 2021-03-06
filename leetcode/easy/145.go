package easy

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int

	// εεΊιε
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}
