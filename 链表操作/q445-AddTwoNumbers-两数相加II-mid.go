package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1 []int
	var s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var head *ListNode
	var carry = 0
	sumval, x, y := 0, 0, 0
	for len(s1) > 0 || len(s2) > 0 {
		if len(s1) > 0 {
			x = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		} else {
			x = 0
		}
		if len(s2) > 0 {
			y = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		} else {
			y = 0
		}
		sumval = x + y + carry
		carry = sumval / 10
		pre := &ListNode{Val: sumval % 10}
		pre.Next = head
		head = pre
	}
	if carry > 0 {
		head = &ListNode{Val: carry, Next: head}
	}
	return head
}
