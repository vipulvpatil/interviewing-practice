package sortableelements

import (
	"fmt"
)

type SortableInts struct {
	values []int
}

func NewSortableInts(v ...int) *SortableInts {
	return &SortableInts{values: v}
}

func (s *SortableInts) Swap(i, j int) {
	temp := s.values[i]
	s.values[i] = s.values[j]
	s.values[j] = temp
}
func (s *SortableInts) Less(i, j int) bool {
	return s.values[i] < s.values[j]
}

func (s *SortableInts) Len() int {
	return len(s.values)
}

func (s *SortableInts) String() string {
	stringifiedValue := ""
	for i, x := range s.values {
		stringifiedValue = fmt.Sprintf("%s%d", stringifiedValue, x)
		if i < len(s.values)-1 {
			stringifiedValue = fmt.Sprintf("%s%s", stringifiedValue, ", ")
		}
	}
	return stringifiedValue
}
