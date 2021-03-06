package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil && t1.Val == t2.Val {
		return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	} else {
		return false
	}
}
