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
    if root.Left != nil {
        if root.Right != nil {
            root.Left.Next = root.Right
        } else {
            root.Left.Next = search(root)
        }
    }
    if root.Right != nil {
        root.Right.Next = search(root)
    }
    connect(root.Right)
    connect(root.Left)
    return root
}

func search(root *Node) *Node {
    for node := root; node != nil; node = node.Next {
        if node.Next != nil {
            if node.Next.Left != nil {
                return node.Next.Left
            } else if node.Next.Right != nil {
                return node.Next.Right
            }
        }
    }
    return nil
}