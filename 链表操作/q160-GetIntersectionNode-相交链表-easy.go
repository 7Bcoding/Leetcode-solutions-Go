package 链表操作

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(pHead1, pHead2 *ListNode) *ListNode {
	nodeMap := map[*ListNode]int{}
	for pHead1 != nil{
		nodeMap[pHead1]++
		pHead1 = pHead1.Next
	}
	for pHead2 != nil{
		if nodeMap[pHead2] > 0{
			return pHead2
		}
		pHead2 = pHead2.Next
	}
	return nil

}