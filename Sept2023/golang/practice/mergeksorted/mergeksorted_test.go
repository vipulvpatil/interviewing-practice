package mergeksorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerging(t *testing.T) {
	t.Run("testing merging", func(t *testing.T) {
		result := MergeKLists(createListOfLinkList([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}))
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, createArray(result))
		result = MergeKLists(createListOfLinkList([][]int{{}}))
		assert.Equal(t, []int{}, createArray(result))
		result = MergeKLists(createListOfLinkList([][]int{}))
		assert.Equal(t, []int{}, createArray(result))
	})
}

func createListOfLinkList(arr [][]int) []*ListNode {
	result := []*ListNode{}
	for i := range arr {
		result = append(result, createLinkList(arr[i]))
	}
	return result
}

func createLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNode{Val: arr[0]}
	curr := root

	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return root
}

func createArray(node *ListNode) []int {
	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
