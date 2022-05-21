package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nodeMap := map[*ListNode]int{}
	for head != nil{
		if nodeMap[head] == 0 {
			nodeMap[head]++
		} else{
			return head
		}
		head = head.Next
	}

	return nil
}
