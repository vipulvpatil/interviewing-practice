package shell

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"

type ShellSorter struct{}

func (s *ShellSorter) Sort(arr sortable.SortableInPlace) {
	n := arr.Len()
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for ; h >= 1; h = h / 3 {
		s.HSort(arr, h)
	}
}

func (s *ShellSorter) HSort(arr sortable.SortableInPlace, h int) {
	for i := h; i < arr.Len(); i++ {
		for j := i; j >= h; j -= h {
			if arr.Less(j, j-h) {
				arr.Swap(j, j-h)
			} else {
				break
			}
		}
	}
}

func NewSorter() sortable.InPlaceSorter {
	return &ShellSorter{}
}
