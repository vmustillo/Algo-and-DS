package tree

import (
	"fmt"
	"bytes"
	"strconv"
)

// PreorderSearch is a function that searches for the target int in a Binary Search Tree by searching the root node, and then the left and right subtrees, respectively
// if target is found, return true, else return false
func PreorderSearch(root *Node, target int) bool {
	var visited []*Node
	var temp *Node

	visited = append(visited, root)

	for len(visited) > 0 && root != nil {
		temp = visited[len(visited)-1]
		if temp.Val == target {
			return true
		}
		visited = visited[0:len(visited)-1]
		if temp.Right != nil {
			visited = append(visited, temp.Right)
		}
		if temp.Left != nil {
			visited = append(visited, temp.Left)
		}
	}
	return false
}

// Preorder traverses the tree in preorder fashion (root --> left --> right)
func Preorder(root *Node) {
	var visited []*Node
	var temp *Node
	var b bytes.Buffer

	visited = append(visited, root)

	for len(visited) > 0 && root != nil {
		temp = visited[len(visited)-1]
		i := strconv.Itoa(temp.Val)
		b.WriteString(i)
		b.WriteString(",")

		visited = visited[0:len(visited)-1]
		if temp.Right != nil {
			visited = append(visited, temp.Right)
		}
		if temp.Left != nil {
			visited = append(visited, temp.Left)
		}
	}
	s := b.String()
	fmt.Println(s[0:len(s)-1])
}