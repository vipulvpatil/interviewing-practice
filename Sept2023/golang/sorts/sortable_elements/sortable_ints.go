package sortableelements

import (
	"fmt"
)

type SortableInts struct {
	Values []int
}

func NewSortableInts(v ...int) *SortableInts {
	return &SortableInts{Values: v}
}

func (s *SortableInts) Swap(i, j int) {
	temp := s.Values[i]
	s.Values[i] = s.Values[j]
	s.Values[j] = temp
}
func (s *SortableInts) Less(i, j int) bool {
	return s.Values[i] < s.Values[j]
}

func (s *SortableInts) Len() int {
	return len(s.Values)
}

func (s *SortableInts) Get(i int) any {
	return s.Values[i]
}

func (s *SortableInts) Set(i int, value any) {
	s.Values[i] = value.(int)
}

func (s *SortableInts) String() string {
	stringifiedValue := ""
	for i, x := range s.Values {
		stringifiedValue = fmt.Sprintf("%s%d", stringifiedValue, x)
		if i < len(s.Values)-1 {
			stringifiedValue = fmt.Sprintf("%s%s", stringifiedValue, ", ")
		}
	}
	return stringifiedValue
}
