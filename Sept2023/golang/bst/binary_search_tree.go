package bst

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinarySearchTree[K constraints.Ordered, V any] struct {
	root *TreeNode[K, V]
}

func (b *BinarySearchTree[K, V]) InOrder() []*TreeNode[K, V] {
	return b.root.inOrder()
}

func (b *BinarySearchTree[K, V]) Insert(key K, value V) {
	b.root = b.root.add(key, value)
}

func (b *BinarySearchTree[K, V]) Print() string {
	str := ""
	nodes := b.InOrder()
	for i, node := range nodes {
		str = fmt.Sprintf("%s%v:%v", str, node.key, node.value)
		if i < len(nodes)-1 {
			str = fmt.Sprintf("%s, ", str)
		}
	}
	return str
}

type TreeNode[K constraints.Ordered, V any] struct {
	key   K
	value V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
}

func (t *TreeNode[K, V]) inOrder() []*TreeNode[K, V] {
	arr := []*TreeNode[K, V]{}
	if t.left != nil {
		arr = append(arr, t.left.inOrder()...)
	}
	arr = append(arr, t)
	if t.right != nil {
		arr = append(arr, t.right.inOrder()...)
	}
	return arr
}

func (t *TreeNode[K, V]) add(key K, value V) *TreeNode[K, V] {
	if t == nil {
		return &TreeNode[K, V]{
			key:   key,
			value: value,
		}
	}

	if key < t.key {
		t.left = t.left.add(key, value)
	} else if key > t.key {
		t.right = t.right.add(key, value)
	} else {
		t.value = value
	}
	return t
}
