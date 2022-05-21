package 链表操作

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

func hasCycle( head *ListNode) bool {
	nodeMap := map[*ListNode]int{}
	for head != nil{
		if nodeMap[head] == 0{
			nodeMap[head]++
		}else{
			return true
		}
		head = head.Next
	}
	return false
}