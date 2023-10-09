package quick3

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type Quick3Sorter struct{}

func (s *Quick3Sorter) Sort(arr sortable.SortableInPlace) {
	Quick3Sort(arr, 0, arr.Len()-1)
}

func Quick3Sort(arr sortable.SortableInPlace, start, end int) {
	if end <= start {
		return
	}
	i := start
	lt := start
	gt := end
	for i <= gt {
		if arr.Less(i, lt) {
			arr.Swap(i, lt)
			i++
			lt++
		} else if arr.Less(lt, i) {
			arr.Swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	Quick3Sort(arr, start, lt-1)
	Quick3Sort(arr, gt+1, end)
}

func NewSorter() sortable.InPlaceSorter {
	return &Quick3Sorter{}
}
