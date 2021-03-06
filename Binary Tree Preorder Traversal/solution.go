package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	nums = append(nums, root.Val)
	nums = append(nums, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)
	return nums
}
