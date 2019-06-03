package main

import (
	"fmt"
	"bst/tree"
)

func main() {
	root := tree.NewNode(25)

	root.Insert(15)
	root.Insert(6)
	root.Insert(4)
	root.Insert(23)
	root.Insert(17)
	root.Insert(16)
	root.Insert(19)
	root.Insert(24)
	root.Insert(28)
	root.Insert(35)
	root.Insert(29)
	root.Insert(42)

	/*
						25
					/		\
				15		     28
			  /	  \			   \
			 6	  23			 35
		   /	 /	\			/  \
		  4		17	24		  29	42
			   /  \
			16     19
	*/

	fmt.Println("Level-order traversal: ")
	tree.BFTraversal(root)
	fmt.Println("\nPreorder traversal: ")
	tree.Preorder(root)
	fmt.Println("\nInorder traversal: ")
	tree.Inorder(root)
	fmt.Println("\nPostorder traversal: ")
	tree.Postorder(root)
}
