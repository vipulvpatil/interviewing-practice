package heap

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type HeapSorter struct{}

func leftChild(n int) int {
	return 2*n + 1
}

func rightChild(n int) int {
	return 2*n + 2
}

func parent(n int) int {
	return (n - 1) / 2
}

func sink(arr sortable.SortableInPlace, i int, size int) {
	if i >= size {
		return
	}
	left := leftChild(i)
	right := rightChild(i)
	next := size
	if left < size && arr.Less(i, left) {
		next = left
	}
	if right < size && arr.Less(i, right) {
		if arr.Less(left, right) {
			next = right
		}
	}

	if next < size {
		arr.Swap(i, next)
		sink(arr, next, size)
	}
}

func swim(arr sortable.SortableInPlace, i int) {
	if i <= 0 {
		return
	}
	p := parent(i)
	if arr.Less(p, i) {
		arr.Swap(i, p)
		swim(arr, p)
	}
}

func (s *HeapSorter) Sort(arr sortable.SortableInPlace) {
	for i := parent(arr.Len() - 1); i >= 0; i-- {
		sink(arr, i, arr.Len())
	}

	for size := arr.Len() - 1; size > 0; size-- {
		arr.Swap(0, size)
		sink(arr, 0, size)
	}
}

func NewSorter() sortable.InPlaceSorter {
	return &HeapSorter{}
}
