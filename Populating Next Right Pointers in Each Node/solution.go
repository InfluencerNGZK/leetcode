/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    for leftMost := root; leftMost.Left != nil; leftMost = leftMost.Left {
        for node := leftMost; node != nil; node = node.Next {
            node.Left.Next = node.Right
            if node.Next != nil {
                node.Right.Next = node.Next.Left
            }
        }
    }
    return root
}