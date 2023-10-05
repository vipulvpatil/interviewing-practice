package selection

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type SelectionSorter struct{}

func (s *SelectionSorter) Sort(arr sortable.SortableInPlace) {
	for i := 0; i < arr.Len(); i++ {
		minIndex := i
		for j := i + 1; j < arr.Len(); j++ {
			if arr.Less(j, minIndex) {
				minIndex = j
			}
		}
		arr.Swap(i, minIndex)
	}
}

func NewSorter() sortable.InPlaceSorter {
	return &SelectionSorter{}
}
