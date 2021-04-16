package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		if fast == slow {
			break
		}
	}
	ptr := head
	for ptr != slow {
		if slow.Next != nil {
			ptr = ptr.Next
			slow = slow.Next
		} else {
			return nil
		}
	}
	return ptr
}
