package tree

import (
	"fmt"
	"bytes"
	"strconv"
)

// BFSearch returns a boolean representing whether the target passed is inside the binary tree passed
// takes a pointer to the root node of a binary search tree and an integer to look for
// If the node passed is not false, it will add the leaf to the slice of leaf nodes
// While the length of the slice is greater than zero, it will iterate through each leaf checking to the see if its value is equal to the target.
// If the value is not equal to the target it appends the left and then right leaf nodes (if they are not null) to the end of the slice.
// After each leaf is checked, if the value has not been found, the slice is shifted to the right to get rid of the leaf that was just checked 
func BFSearch(root *Node, target int) bool {
	var leavesOnCurrLevel []*Node

	if root == nil {
		return false
	} 
	
	leavesOnCurrLevel = append(leavesOnCurrLevel, root)

	for len(leavesOnCurrLevel) > 0 {
		for _, it := range leavesOnCurrLevel {
			if it.Val == target {
				return true
			}
			if it.Left != nil {
				leavesOnCurrLevel = append(leavesOnCurrLevel, it.Left)
			}
			if it.Right != nil {
				leavesOnCurrLevel = append(leavesOnCurrLevel, it.Right)
			}
			leavesOnCurrLevel = leavesOnCurrLevel[1:]
		}
	}
	return false
}

// BFTraversal traverses through the tree level by level, printing out nodes
func BFTraversal(root *Node) {
	var leavesOnCurrLevel []*Node
	var b bytes.Buffer

	leavesOnCurrLevel = append(leavesOnCurrLevel, root)

	for len(leavesOnCurrLevel) > 0 && root != nil {
		for _, it := range leavesOnCurrLevel {
			i := strconv.Itoa(it.Val)
			b.WriteString(i)
			b.WriteString(",")

			if it.Left != nil {
				leavesOnCurrLevel = append(leavesOnCurrLevel, it.Left)
			}
			if it.Right != nil {
				leavesOnCurrLevel = append(leavesOnCurrLevel, it.Right)
			}
			leavesOnCurrLevel = leavesOnCurrLevel[1:]
		}
	}

	s := b.String()
	fmt.Println(s[0:len(s)-1])
}