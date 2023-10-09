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
	j := end
	for {
		for i <= end && arr.Less(i, pivot) {
			i++
		}
		for j >= start+1 && arr.Less(pivot, j) {
			j--
		}

		if i < j {
			arr.Swap(i, j)
		} else {
			break
		}
	}
	arr.Swap(pivot, j)
	QuickSort(arr, start, j-1)
	QuickSort(arr, j+1, end)
}

func NewSorter() sortable.InPlaceSorter {
	return &QuickSorter{}
}
