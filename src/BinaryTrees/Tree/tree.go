package tree

import "fmt"

// Node structure will be used to create binary search tree
// Each leaf Node will contain a value and pointer to its left and right child leaf nodes
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// NewNode creates a node with the value entered and nil Left and Right values
func NewNode(val int) *Node {
	return &Node{Val: val}
}

// Insert inserts a leaf node into a binary tree
// val parameter contains the value of the leaf node being inserted
func (n *Node) Insert(val int) {
	if val <= n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else if val > n.Val {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

// PrintNode outputs the value of the current node and its child nodes
func PrintNode(n *Node) {
	fmt.Print("Value: ", n.Val, "\n")
	if n.Left != nil {
		fmt.Println("Left: ", n.Left.Val)
	}
	if n.Right != nil {
		fmt.Println("Right: ", n.Right.Val)
	}
}