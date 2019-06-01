package main

import "fmt"

// Node structure will be used to create binary search tree
// Each leaf Node will contain a value and pointer to its left and right child leaf nodes
type Node struct {
	Val   int
	Left *Node
	Right *Node
}

func inorder(root *Node) {
	var stack []*Node
	var temp *Node

	temp = root

	for len(stack) > 0 || temp != nil {
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}

		temp = stack[len(stack)-1]
		fmt.Println(temp.Val)
		stack = stack[0:len(stack)-1]

		temp = temp.Right

		// if temp.Left != nil {
		// 	tempStack = append(tempStack, temp.Left)
		// }
	}
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

	inorder(&root)
}