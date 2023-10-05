package merge

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
)

type MergeSorter struct{}

func (s *MergeSorter) Sort(arr sortable.SortableOutOfPlace) {
	for sz := 1; sz < arr.Len(); sz = sz * 2 {
		for start := 0; start < arr.Len(); start += sz * 2 {
			end := start + sz*2 - 1
			if end > arr.Len()-1 {
				end = arr.Len() - 1
			}
			mid := start + sz - 1
			MergeSortBottomUp(arr, start, mid, end)
		}
	}
	// OR
	// MergeSort(arr, 0, arr.Len()-1)
}

func MergeSort(arr sortable.SortableOutOfPlace, start, end int) {
	if end-start < 1 {
		return
	}
	mid := start + (end-start)/2
	tempArray := []any{}
	MergeSort(arr, start, mid)
	MergeSort(arr, mid+1, end)
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if arr.Less(i, j) {
			tempArray = append(tempArray, arr.Get(i))
			i++
		} else {
			tempArray = append(tempArray, arr.Get(j))
			j++
		}
	}
	if i > mid {
		for j <= end {
			tempArray = append(tempArray, arr.Get(j))
			j++
		}
	}
	if j > end {
		for i <= mid {
			tempArray = append(tempArray, arr.Get(i))
			i++
		}
	}
	for k := range tempArray {
		arr.Set(k+start, tempArray[k])
	}
}

func MergeSortBottomUp(arr sortable.SortableOutOfPlace, start, mid, end int) {
	if end-start < 1 {
		return
	}

	tempArray := []any{}
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if arr.Less(i, j) {
			tempArray = append(tempArray, arr.Get(i))
			i++
		} else {
			tempArray = append(tempArray, arr.Get(j))
			j++
		}
	}
	if i > mid {
		for j <= end {
			tempArray = append(tempArray, arr.Get(j))
			j++
		}
	}
	if j > end {
		for i <= mid {
			tempArray = append(tempArray, arr.Get(i))
			i++
		}
	}
	for k := range tempArray {
		arr.Set(k+start, tempArray[k])
	}
}

func NewSorter() sortable.OutOfPlaceSorter {
	return &MergeSorter{}
}
