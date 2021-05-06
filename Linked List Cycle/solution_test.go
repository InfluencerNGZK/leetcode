package main

import (
	"fmt"
	"testing"
)

func InitNode(val int, next *ListNode) (ret *ListNode){
	ret = new(ListNode)
	ret.Val = val
	ret.Next = next

	return ret
}

func Test(t *testing.T) {
	l4 := InitNode(-4, nil)
	l3 := InitNode(0, l4)
	l2 := InitNode(2, l3)
	l1 := InitNode(3, l2)
	l4.Next = l2
	fmt.Println(hasCycle(l1))
}
