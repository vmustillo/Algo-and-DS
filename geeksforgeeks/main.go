// Given a linked list which is sorted, how will you insert in sorted way

package main

import (
	"github.com/vmustillo/algo-and-ds/geeksforgeeks/linkedlist"
)

func main() {
	head := linkedlist.InsertSorted(nil, 1)
	head = linkedlist.InsertSorted(head, 4)
	head = linkedlist.InsertSorted(head, 6)
	head = linkedlist.InsertSorted(head, 8)
	head = linkedlist.InsertSorted(head, 12)
	head = linkedlist.InsertSorted(head, 15)

	head = linkedlist.InsertSorted(head, 7)
	head = linkedlist.InsertSorted(head, 0)
	head = linkedlist.InsertSorted(head, -1)

	linkedlist.PrintList(head)

	linkedlist.DeleteNode(head, head)
	linkedlist.PrintList(head)

	n := &linkedlist.Node{2, nil}
	linkedlist.DeleteNode(n, n)
	linkedlist.PrintList(n)
}