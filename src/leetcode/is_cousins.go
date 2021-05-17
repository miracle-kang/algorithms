package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 993. 二叉树的堂兄弟节点
// 在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
// 如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
// 我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
// 只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
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
