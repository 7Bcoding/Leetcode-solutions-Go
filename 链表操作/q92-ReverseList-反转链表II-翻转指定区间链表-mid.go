package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || left == right {
		return head
	}
	revHead := head
	revHeadPre := revHead
	i := 1
	for i < left {
		revHeadPre = revHead
		revHead = revHead.Next
		i++
	}
	// 待翻转链表原来的头
	var pre *ListNode
	var tmp *ListNode
	curr := revHead
	j := left
	for curr != nil && j <= right {
		tmp = curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
		j++
	}

	revHead.Next = curr
	if left != 1{
		revHeadPre.Next = pre
	} else{
		head = pre
	}
	return head
}
