package tree

import (
	"fmt"
	"bytes"
	"strconv"
)

// Postorder traverses the tree in postorder fashion (left --> right --> root)
func Postorder(root *Node) {
	var resultStack []*Node
	var tempStack []*Node
	var temp *Node
	var b bytes.Buffer

	tempStack = append(tempStack, root)

	for len(tempStack) > 0 && root != nil {
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
		s := strconv.Itoa(resultStack[i].Val)
		b.WriteString(s)
		b.WriteString(",")
	}
	str := b.String()
	fmt.Println(str[0:len(str)-1])
}

func PostorderRecursive(root *Node) {
	if root == nil {
		return
	}

	PostorderRecursive(root.Left)
	PostorderRecursive(root.Right)
	fmt.Printf("%d,", root.Val)
}