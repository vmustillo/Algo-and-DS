package main

import (
	"bst/tree"
	"fmt"
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

	subtreeRoot := tree.NewNode(35)
	subtreeRoot.Insert(29)
	subtreeRoot.Insert(42)

	/*
						   SUBTREE

						   35
						/		\
					29			42

		/*                 TREE

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

	// fmt.Println("Level-order traversal ")
	// tree.BFTraversal(root)
	// fmt.Println("\nPreorder traversal (Iteratively then Recursively): ")
	// tree.Preorder(root)
	// tree.PreorderRecursive(root)
	// fmt.Println("\nInorder traversal (Iteratively then Recursively): ")
	// tree.Inorder(root)
	// tree.InorderRecursive(root)
	// fmt.Println("\nPostorder traversal (Iteratively then Recursively): ")
	// tree.Postorder(root)
	// tree.PostorderRecursive(root)

	if tree.IsSubtree(root, subtreeRoot) == true {
		fmt.Println("true!")
	} else {
		fmt.Println("false!")
	}
}
