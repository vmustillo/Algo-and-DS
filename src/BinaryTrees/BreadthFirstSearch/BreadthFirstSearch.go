package main

import "fmt"

// TreeNode structure will be used to create binary search tree
// Each leaf Node will contain a value and pointer to its left and right child leaf nodes
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

/* 
	bfSearch returns a boolean representing whether the target passed is inside the binary tree passed
	
	takes a pointer to the root node of a binary search tree and an integer to look for
	If the node passed is not false, it will add the leaf to the slice of leaf nodes

	While the length of the slice is greater than zero, it will iterate through each leaf checking to the see if its value is equal to the target.
	If the value is not equal to the target it appends the left and then right leaf nodes (if they are not null) to the end of the slice.
	After each leaf is checked, if the value has not been found, the slice is shifted to the right to get rid of the leaf that was just checked 
*/
func bfSearch(root *TreeNode, target int) bool {
	var leavesOnCurrLevel []*TreeNode

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

func main() {
	leaf1 := TreeNode{2, nil, nil}
	leaf2 := TreeNode{5, &leaf1, nil}
	leaf3 := TreeNode{9, nil, nil}
	leaf4 := TreeNode{15, nil, nil}
	leaf5 := TreeNode{12, &leaf3, &leaf4}
	root := TreeNode{7, &leaf2, &leaf5}

	/*
						7
					/		\
				5				12
			/				/		\
		2				9				15
		
	
	*/

	target := 6
	if bfSearch(&root, target) {
		fmt.Println(target, "is in the tree! :)")
	} else {
		fmt.Println(target, "is not in the tree :(")
	}
}