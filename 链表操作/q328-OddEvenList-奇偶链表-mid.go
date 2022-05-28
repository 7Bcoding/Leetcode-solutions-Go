package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyOdd := ListNode{-1, nil}
	dummyEven := ListNode{-1, nil}
	odd, even := head, head.Next
	dummyOdd.Next = odd
	dummyEven.Next = even

	for odd != nil && even != nil {
		if odd.Next.Next != nil {
			odd.Next = odd.Next.Next
			odd = odd.Next
		} else {
			odd.Next = dummyEven.Next
			even.Next = nil
			break
		}
		if even.Next.Next != nil {
			even.Next = even.Next.Next
			even = even.Next
		} else {
			odd.Next = dummyEven.Next
			even.Next = nil
			break
		}
	}
	return dummyOdd.Next
}
