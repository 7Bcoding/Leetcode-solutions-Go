package main

import (
	. "fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre, p1, p2 := head, head, head
	dist := 1
	for p2.Next != nil && dist < n{
		p2 = p2.Next
		dist++
	}
	for p2.Next != nil{
		pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	node := p1
	pre.Next = node.Next
	if node == head{
		return node.Next
	}
	return head
}
func main()  {
	var head *ListNode
	head = &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						nil}}}}}
	var node *ListNode
	node = head
	for node != nil{
		Print("",node.Val)
		Print("->")
		node = node.Next
	}
	Print("nil")
	newHead := removeNthFromEnd(head, 2)
	Println("")
	node = newHead
	for node != nil{
		Print("",node.Val)
		Print("->")
		node = node.Next
	}
	Print("nil")
}

