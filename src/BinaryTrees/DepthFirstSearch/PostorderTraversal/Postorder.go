package main

import (
	"fmt"
	"os"
)

// Node structure will be used to create binary search tree
// Each leaf Node will contain a value and pointer to its left and right child leaf nodes
type Node struct {
	Val   int
	Left *Node
	Right *Node
}

func postorder(root *Node) {
	var resultStack []*Node
	var tempStack []*Node
	var temp *Node

	if root == nil {
		os.Exit(1)
	}

	tempStack = append(tempStack, root)

	for len(tempStack) > 0 {
		temp = tempStack[len(tempStack)-1]
		resultStack = append(resultStack, tempStack[len(tempStack)-1])
		tempStack = tempStack[0:len(tempStack)-1]
		if temp.Left != nil {
			tempStack = append(tempStack, temp.Left)
		}
		if temp.Right != nil {
			tempStack = append(tempStack, temp.Right)
		}
	}

	for i := len(resultStack)-1; i >= 0; i-- {
		fmt.Println(resultStack[i].Val)
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

	postorder(&root)
}