package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := 0
	rightCount := 0
	leftCount = 1 + maxDepth(root.Left)
	rightCount = 1 + maxDepth(root.Right)
	if leftCount > rightCount {
		return leftCount
	} else {
		return rightCount
	}
}
