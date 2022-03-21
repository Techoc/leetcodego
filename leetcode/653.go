package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	//建立一个Hash表
	set := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		//如果哈希表中是否存在 k - Val 则返回true
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		//将Val添加进Hash表
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
