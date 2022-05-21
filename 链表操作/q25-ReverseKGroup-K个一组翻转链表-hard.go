package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{-1, nil}
	dummy.Next = head
	pre := dummy
	for head != nil{
		tail := pre
		for i:=0; i<k; i++{
			tail = tail.Next
			if tail == nil{
				return dummy.Next
			}
		}
		post := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = post
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode)  {
	start := tail
	end := head
	tail.Next = nil
	pre, curr := &ListNode{}, head
	for curr != nil{
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return start, end
}
