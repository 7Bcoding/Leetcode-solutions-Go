package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil{
		tmp := curr.Next
		// 结构体指针类型不支持使用:=的方式传递值
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre

}
