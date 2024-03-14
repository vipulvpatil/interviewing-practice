package mergeksorted

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	pq := new(minPQ)
	for _, l := range lists {
		if l != nil {
			heap.Push(pq, l)
		}
	}
	var result *ListNode
	var node *ListNode
	for pq.Len() > 0 {
		minNode := heap.Pop(pq).(*ListNode)
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

type minPQ []*ListNode

func (pq minPQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq minPQ) Len() int {
	return len(pq)
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minPQ) Push(n any) {
	*pq = append(*pq, n.(*ListNode))
}

func (pq *minPQ) Pop() any {
	t := *pq
	l := len(t)
	result := t[l-1]
	*pq = t[0 : l-1]
	return result
}
