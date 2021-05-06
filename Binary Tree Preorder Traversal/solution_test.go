package main

import (
	"fmt"
	"testing"
)

func InitNode(val int, left *TreeNode, right *TreeNode) (ret *TreeNode){
	ret = new(TreeNode)
	ret.Val = val
	ret.Left = left
	ret.Right = right

	return ret
}

// RUN COMMAND "go test -test.run <function name>"
func Test(t *testing.T) {
	l3 := InitNode(3, nil, nil)
	l2 := InitNode(2, l3, nil)
	l1 := InitNode(1, nil, l2)
	fmt.Println(preorderTraversal(l1))
}
