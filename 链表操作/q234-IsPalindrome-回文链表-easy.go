package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i := 0; i <= n/2; i++ {
		if vals[i] != vals[n-1-i] {
			return false
		}
	}
	return true
}
