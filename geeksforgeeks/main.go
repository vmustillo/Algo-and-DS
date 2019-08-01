// Given a linked list which is sorted, how will you insert in sorted way

package main

import (
	"fmt"
	"strings"

	"github.com/vmustillo/algo-and-ds/geeksforgeeks/linkedlist"
)

func main() {
	empty := &linkedlist.Node{}
	head := linkedlist.InsertSorted(empty, 1)
	head = linkedlist.InsertSorted(head, 4)
	head = linkedlist.InsertSorted(head, 6)
	head = linkedlist.InsertSorted(head, 8)
	head = linkedlist.InsertSorted(head, 12)
	head = linkedlist.InsertSorted(head, 15)

	head = linkedlist.InsertSorted(head, 7)
	head = linkedlist.InsertSorted(head, 0)
	head = linkedlist.InsertSorted(head, -1)
	var sb strings.Builder

	for head != nil {
		fmt.Fprintf(&sb, "%d --> ", head.Val)
		head = head.Next
	}
	s := sb.String()
	fmt.Println(s[:len(s) - 5])
}