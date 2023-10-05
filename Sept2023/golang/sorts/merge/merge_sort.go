package merge

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type MergeSorter struct{}

func (s *MergeSorter) Sort(arr sortable.SortableOutOfPlace) {
	for i := 1; i < arr.Len(); i++ {
		for j := i; j > 0; j-- {
			if arr.Less(j, j-1) {
				// arr.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}

// func MergeSort(arr)

func NewSorter() sortable.OutOfPlaceSorter {
	return &MergeSorter{}
}
