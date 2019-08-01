package linkedlist

type Node struct {
	Val int
	Next *Node
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