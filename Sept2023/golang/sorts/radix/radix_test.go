package radix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BaseRadixSort(t *testing.T) {
	arr := []string{"A", "Z", "a", "c", "3", "", "g", "H", "T", "t", "5", "å"}
	Sort(arr, 0, 0, len(arr)-1)
	fmt.Println(arr)
	assert.Equal(t, 12, len(arr))
	assert.Fail(t, "testing")
}

// func Test_LSDRadixSort(t *testing.T) {
// 	arr := []string{"A", "Z", "a", "c", "3", "g", "H", "T", "t", "5"}
// 	Sort(arr, 0, 0, len(arr)-1)
// 	fmt.Println(arr)
// 	assert.Fail(t, "testing")
// }
