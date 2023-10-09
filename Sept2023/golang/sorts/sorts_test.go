package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/insertion"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/merge"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/quick"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/quick3"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/selection"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/shell"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
	sortableelements "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable_elements"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		name             string
		input            []int
		expectedOutput   []int
		inPlaceSorter    sortable.InPlaceSorter
		outOfPlaceSorter sortable.OutOfPlaceSorter
	}{
		{
			name:           "Selection sort",
			input:          []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:  selection.NewSorter(),
		},
		{
			name:           "Insertion sort",
			input:          []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:  insertion.NewSorter(),
		},
		{
			name:           "Shell sort",
			input:          []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:  shell.NewSorter(),
		},
		{
			name:             "Merge sort",
			input:            []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:    nil,
			outOfPlaceSorter: merge.NewSorter(),
		},
		{
			name:           "Quick sort",
			input:          []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:  quick.NewSorter(),
		},
		{
			name:           "3 way Quick sort",
			input:          []int{8, 4, 2, 6, 3, 7, 1, 9, 5},
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			inPlaceSorter:  quick3.NewSorter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := sortableelements.NewSortableInts(tt.input...)
			if tt.inPlaceSorter != nil {
				tt.inPlaceSorter.Sort(arr)
			} else {
				tt.outOfPlaceSorter.Sort(arr)
			}
			assert.Equal(t, tt.expectedOutput, arr.Values)
		})
	}

}
