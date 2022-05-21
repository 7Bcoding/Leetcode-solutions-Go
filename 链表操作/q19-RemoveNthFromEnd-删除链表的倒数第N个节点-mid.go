package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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
