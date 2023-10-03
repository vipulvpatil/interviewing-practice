package sorts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/insertion"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/selection"
	sortableelements "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/sorts/sortable_elements"
)

func Test_SelectionSort(t *testing.T) {
	t.Run("Selection sort works", func(t *testing.T) {
		arr := sortableelements.NewSortableInts(8, 4, 2, 6, 3, 7, 1, 9, 5)
		selection.NewSorter().Sort(arr)
		output := fmt.Sprint(arr)
		assert.Equal(t, "1, 2, 3, 4, 5, 6, 7, 8, 9", output)
	})
}

func Test_InsertionSort(t *testing.T) {
	t.Run("Insertion sort works", func(t *testing.T) {
		arr := sortableelements.NewSortableInts(8, 4, 2, 6, 3, 7, 1, 9, 5)
		insertion.NewSorter().Sort(arr)
		output := fmt.Sprint(arr)
		assert.Equal(t, "1, 2, 3, 4, 5, 6, 7, 8, 9", output)
	})
}
