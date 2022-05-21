package 链表操作


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, &ListNode{-1, nil}}
	head := dummy
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			head.Next = l1
			l1 = l1.Next
		} else{
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil{
		head.Next = l1
	} else{
		head.Next = l2
	}
	return dummy.Next
}