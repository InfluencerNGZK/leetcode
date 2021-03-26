/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }
    p1 := head
    p2 := head
    for p2.Next != nil && p2.Next.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    
    var prev *ListNode
    curr := p1.Next
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    
    for prev != nil {
        if head.Val != prev.Val {
            return false
        }
        head = head.Next
        prev = prev.Next
    }
    return true
}