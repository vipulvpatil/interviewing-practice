package insertion

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type InsertionSorter struct{}

func (s InsertionSorter) Sort(arr sortable.Sortable) {
	for i := 1; i < arr.Len(); i++ {
		for j := i; j > 0; j-- {
			if arr.Less(j, j-1) {
				arr.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}

func NewSorter() sortable.Sorter {
	return InsertionSorter{}
}
