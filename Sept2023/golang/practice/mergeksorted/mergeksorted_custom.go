package mergeksorted

import "container/heap"

func MergeKListsCustom(lists []*ListNode) *ListNode {
	pq := new(minPQ)
	for _, l := range lists {
		if l != nil {
			Push(pq, l)
		}
	}
	var result *ListNode
	var node *ListNode
	for pq.Len() > 0 {
		minNode := Pop(pq).(*ListNode)
		newNode := &ListNode{
			Val: minNode.Val,
		}
		if result == nil {
			result = newNode
			node = newNode
		} else {
			node.Next = newNode
			node = newNode
		}
		if minNode.Next != nil {
			heap.Push(pq, minNode.Next)
		}
	}

	return result
}
