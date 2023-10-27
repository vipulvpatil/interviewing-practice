package radix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BaseRadixSort(t *testing.T) {
	arr := []string{"A", "Z", "a", "c", "3", "", "g", "H", "T", "t", "5", "å"}
	expected := []string{"", "å", "3", "5", "A", "H", "T", "Z", "a", "c", "g", "t"}
	Sort(arr, 0, 0, len(arr)-1)
	assert.Equal(t, len(expected), len(arr))
	assert.Equal(t, expected, arr)
}

func Test_LSDRadixSort(t *testing.T) {
	arr := []string{"A", "Z", "bat", "cat", "cataract", "3", "", "g", "bad", "H", "T", "t", "5", "cast", "å"}
	expected := []string{"", "å", "3", "5", "A", "H", "T", "Z", "bad", "bat", "cast", "cat", "cataract", "g", "t"}
	LSDSort(arr)
	assert.Equal(t, len(expected), len(arr))
	assert.Equal(t, expected, arr)
}

func Test_MSDRadixSort(t *testing.T) {
	arr := []string{"Alpha", "Zeta", "bat", "cat", "cataract", "3", "", "g", "bad", "Has", "Test", "åƒ∑5", "t", "0500", "cast", "åƒ∑"}
	expected := []string{"", "åƒ∑", "åƒ∑5", "0500", "3", "Alpha", "Has", "Test", "Zeta", "bad", "bat", "cast", "cat", "cataract", "g", "t"}
	NewMSDSort(arr)
	assert.Equal(t, len(expected), len(arr))
	assert.Equal(t, expected, arr)
}

// func Test_ThreeWayQuickSort(t *testing.T) {
// 	arr := []string{"Alpha", "Zeta", "bat", "cat", "cataract", "3", "", "g", "bad", "Has", "Test", "åƒ∑5", "t", "0500", "cast", "åƒ∑"}
// 	expected := []string{"", "åƒ∑", "åƒ∑5", "0500", "3", "Alpha", "Has", "Test", "Zeta", "bad", "bat", "cast", "cat", "cataract", "g", "t"}
// 	ThreeWayQuickSort(arr)
// 	assert.Equal(t, len(expected), len(arr))
// 	assert.Equal(t, expected, arr)
// }
