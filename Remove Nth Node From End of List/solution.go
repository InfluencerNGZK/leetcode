/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fake := &ListNode{0, head}
    pA := fake
    pB := fake
    for i := 1; i < n+2; i++ {
        pA = pA.Next
    }
    for pA != nil {
        pA = pA.Next
        pB = pB.Next
    }
    pB.Next = pB.Next.Next
    return fake.Next
}