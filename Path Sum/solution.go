package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	} else {
		return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
	}
}
