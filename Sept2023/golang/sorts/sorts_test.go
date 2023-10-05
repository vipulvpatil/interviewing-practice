package sorts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/insertion"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/selection"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/shell"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable"
	sortableelements "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable_elements"
)

func Test_Sort(t *testing.T) {
	input := []int{8, 4, 2, 6, 3, 7, 1, 9, 5}
	expectedOutput := "1, 2, 3, 4, 5, 6, 7, 8, 9"
	tests := []struct {
		name   string
		sorter sortable.Sorter
	}{
		{
			name:   "Selection sort",
			sorter: selection.NewSorter(),
		},
		{
			name:   "Insertion sort",
			sorter: insertion.NewSorter(),
		},
		{
			name:   "Shell sort",
			sorter: shell.NewSorter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := sortableelements.NewSortableInts(input...)
			tt.sorter.Sort(arr)
			output := fmt.Sprint(arr)
			assert.Equal(t, expectedOutput, output)
		})
	}

}
