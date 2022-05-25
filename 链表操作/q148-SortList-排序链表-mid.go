package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sortTwoList(head, nil)
}

func sortTwoList(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sortTwoList(head, mid), sortTwoList(mid, tail))
}

func merge(head *ListNode, tail *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	tempNode, pHead1, pHead2 := dummy, head, tail
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			tempNode.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			tempNode.Next = pHead2
			pHead2 = pHead2.Next
		}
		tempNode = tempNode.Next
	}
	if pHead1 != nil {
		tempNode.Next = pHead1
	}
	if pHead2 != nil {
		tempNode.Next = pHead2
	}
	return dummy.Next
}
