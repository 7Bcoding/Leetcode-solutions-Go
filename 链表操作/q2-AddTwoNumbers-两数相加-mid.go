package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addtwonumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	head := dummy
	var carry int
	var x int
	var y int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sumVal := x + y + carry
		remain := sumVal % 10
		carry = sumVal / 10
		head.Next = &ListNode{remain, nil}
		head = head.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry >= 1 {
		head.Next = &ListNode{carry, nil}
	}
	return dummy.Next
}
