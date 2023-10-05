package quick

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type QuickSorter struct{}

func (s *QuickSorter) Sort(arr sortable.SortableInPlace) {
	QuickSort(arr, 0, arr.Len()-1)
}

func QuickSort(arr sortable.SortableInPlace, start, end int) {
	if end <= start {
		return
	}
	pivot := start
	i := start + 1
	j := start + 1
	for j <= end {
		if arr.Less(j, pivot) {
			if i != j {
				arr.Swap(i, j)
			}
			i++
		}
		j++
	}
	arr.Swap(pivot, i-1)
	QuickSort(arr, start, i-2)
	QuickSort(arr, i, end)
}

func NewSorter() sortable.InPlaceSorter {
	return &QuickSorter{}
}
