package main

import "fmt"

// Node structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// function determines if a Linked List contains a cycle
// a cycle causes an infinite loop in a linked list
/*
	This function consists of two pointers, a slow and a fast pointer
	The slow pointer moves through the list node by node
	The fast pointer starts at the second element (if there is one) and moves twice as fast as the slow pointer (skips one node every move)
	If there is an infinite cycle, and one pointer is moving twice as fast as the other pointer, eventually these pointers will land on the same node
	The function checks if two pointers are equal at any point in the traversal and if they are, there is a cycle. If not, eventually the fast pointer will reach nil and there will be no cycle
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var slow *ListNode
	var fast *ListNode
	slow = head
	fast = head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func reverseList(head *ListNode) *ListNode {
	newHead := head
	var prev *ListNode
	for newHead != nil {
		temp := newHead.Next
		newHead.Next = prev
		prev = newHead
		newHead = temp
	}
	return prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("NULL")
}

func main() {
	first := &ListNode{
		Val: 3,
	}
	second := &ListNode{
		Val: 2,
	}
	third := &ListNode{
		Val: 0,
	}
	fourth := &ListNode{
		Val: -4,
	}
	first.Next = second
	second.Val = 2
	second.Next = third
	third.Val = 0
	third.Next = fourth
	fourth.Val = -4
	fourth.Next = second
	fmt.Println("List: ")
	fmt.Println("3 -> 2 -> 0 -> -4")
	fmt.Println("     ^          |")
	fmt.Println("     |	        |")
	fmt.Println("      <- - - - ")
	if hasCycle(first) {
		fmt.Println("Cycle detected")
	} else {
		fmt.Println("No cycle detected")
	}
	fourth.Next = nil
	fmt.Println("New list: ")
	printList(first)
	fmt.Println()
	fmt.Println("Reversed: ")
	printList(reverseList(first))
}
