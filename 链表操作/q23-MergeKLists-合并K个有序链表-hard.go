package 链表操作

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/** 解法1：循环两两合并
 ** 将数组中的链表两两合并，合并后的新链表头和下一个链表进行合并，
 ** 即共进行K-1次

**/
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, curr *ListNode
	k := len(lists)
	for i:=0; i<k; i++{
		if i==0{
			pre = lists[i]
			continue
		}
		curr = lists[i]
		pre = mergeLists(pre, curr)
	}
	return pre
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

/** 解法2：小顶堆
 ** 所有链表头结点加入堆，依次弹出，并获取下一个节点，再push进堆
 ** 弹出的节点，组成新的有序链表

**/

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListHeap(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}

