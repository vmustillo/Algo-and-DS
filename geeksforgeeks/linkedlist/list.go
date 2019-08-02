package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	Val int
	Next *Node
}

func PrintList(head *Node) {
	var sb strings.Builder

	for head != nil {
		fmt.Fprintf(&sb, "%d --> ", head.Val)
		head = head.Next
	}
	s := sb.String()
	fmt.Println(s[:len(s) - 5])
}

// InsertSorted inserts num into a sorted list. Returns the head of the list
func InsertSorted(head *Node, num int) *Node {
	if head == nil {
		h := &Node{
			num,
			nil,
		}

		return h
	}

	if num < head.Val {
		h := &Node{
			num,
			head,
		}

		return h
	}

	follower := head
	tempHead := head.Next

	for tempHead != nil && num > tempHead.Val {
		tempHead = tempHead.Next
		follower = follower.Next
	}

	n := &Node{
		num,
		nil,
	}

	follower.Next = n

	if tempHead != nil {
		n.Next = tempHead
	}

	return head
}

func DeleteNode(head *Node, del *Node) {
	if head == nil {
		return
	}

	if head != nil && head == del {
		if head.Next != nil {
			n := head.Next
			head.Val = n.Val
			head.Next = n.Next.Next
			n = &Node{}
		} else {
			fmt.Println("Cannot delete only node")
		}
		return
	}

	follower := head
	head = head.Next

	for head != nil {
		if head == del {
			follower.Next = head.Next
			head = &Node{}

			return
		}
		head = head.Next
		follower = follower.Next
	}
}