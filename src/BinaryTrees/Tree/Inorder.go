package tree

import (
	"fmt"
	"bytes"
	"strconv"
)

// Inorder traverses the tree in inorder fashion (left --> root --> right)
func Inorder(root *Node) {
	var stack []*Node
	var temp *Node
	var b bytes.Buffer

	temp = root

	for len(stack) > 0 || temp != nil {
		for temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		}

		temp = stack[len(stack)-1]
		i := strconv.Itoa(temp.Val)
		b.WriteString(i)
		b.WriteString(",")
		stack = stack[0:len(stack)-1]

		temp = temp.Right
	}
	s := b.String()
	fmt.Println(s[0:len(s)-1])
}