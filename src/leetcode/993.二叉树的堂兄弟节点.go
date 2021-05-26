/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var xp, yp *TreeNode
	var xd, yd int
	xf, yf := false, false

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil || (xf && yf) {
			return
		}
		if node.Val == x {
			xp, xd, xf = parent, depth, true
		} else if node.Val == y {
			yp, yd, yf = parent, depth, true
		}
		dfs(node.Left, node, depth+1)
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xf && yf && xd == yd && xp != yp
}

// @lc code=end

