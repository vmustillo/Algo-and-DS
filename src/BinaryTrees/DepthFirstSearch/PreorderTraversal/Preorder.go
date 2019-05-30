package main

import "fmt"

// Node structure will be used to create binary search tree
// Each leaf Node will contain a value and pointer to its left and right child leaf nodes
type Node struct {
	Val   int
	Left *Node
	Right *Node
}

// preorder is a function that searches for the target int in a Binary Search Tree by searching the root node, and then the left and right subtrees, respectively
// if target is found, return true, else return false
func preorderSearch(root *Node, target int) bool {
	var visited []*Node
	var temp *Node

	if root == nil {
		return false 
	}

	visited = append(visited, root)

	for len(visited) > 0 {
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

func main() {
	leaf1 := Node{2, nil, nil}
	leaf2 := Node{5, &leaf1, nil}
	leaf3 := Node{9, nil, nil}
	leaf4 := Node{15, nil, nil}
	leaf5 := Node{12, &leaf3, &leaf4}
	root := Node{7, &leaf2, &leaf5}

	/*
						7
					/		\
				5				12
			/				/		\
		2				9				15
		
	
	*/

	target := 6
	if preorderSearch(&root, target) {
		fmt.Println(target, "is in the tree! :)")
	} else {
		fmt.Println(target, "is not in the tree :(")
	}
}