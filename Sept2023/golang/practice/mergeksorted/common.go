package mergeksorted

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
	fmt.Printf("pushing %+v, %d\n", *pq, n)
	*pq = append(*pq, n.(*ListNode))
}

func (pq *minPQ) Pop() any {
	t := *pq
	l := len(t)
	result := t[l-1]
	*pq = t[0 : l-1]
	fmt.Printf("popping %+v, %d\n", *pq, result.Val)
	return result
}

func (pq minPQ) String() string {
	str := ""
	for _, elem := range pq {
		str = fmt.Sprintf("%s %v", str, elem)
	}
	return str
}
