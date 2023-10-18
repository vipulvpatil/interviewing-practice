package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
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

func (b *BinarySearchTree[K, V]) Delete(key K) {
	b.root = b.root.del(key)
}

func (b *BinarySearchTree[K, V]) Search(key K) *TreeNode[K, V] {
	return b.root.search(key)
}

func (b *BinarySearchTree[K, V]) Min() *TreeNode[K, V] {
	return b.root.min()
}

func (b *BinarySearchTree[K, V]) Max() *TreeNode[K, V] {
	return b.root.max()
}

func (b *BinarySearchTree[K, V]) Floor(key K) *TreeNode[K, V] {
	return b.root.floor(key)
}

func (b *BinarySearchTree[K, V]) Ceil(key K) *TreeNode[K, V] {
	return b.root.ceil(key)
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

func PrettyPrint(b *BinarySearchTree[string, string]) {
	levelIndex := make(map[TreeNode[string, string]]int)
	levelIndex[*b.root] = 0
	for _, node := range b.BFS() {
		if node.left != nil {
			levelIndex[*node.left] = levelIndex[*node] + 1
		}
		if node.right != nil {
			levelIndex[*node.right] = levelIndex[*node] + 1
		}
	}
	fmt.Println("printing tree")
	currentLevel := 0
	for _, node := range b.BFS() {
		if levelIndex[*node] != currentLevel {
			currentLevel++
			fmt.Println()
		}

		fmt.Print(" - ")
		if node.left != nil {
			fmt.Printf(" %s<", node.left.value)
		} else {
			fmt.Print(" ")
		}
		fmt.Printf("%s", node.value)
		if node.right != nil {
			fmt.Printf(">%s ", node.right.value)
		}
	}
	fmt.Println()
}

func (b *BinarySearchTree[K, V]) BFS() []*TreeNode[K, V] {
	arr := []*TreeNode[K, V]{}
	q := queue.Queue[TreeNode[K, V]]{}
	q.Enqueue(*b.root)
	for !q.IsEmpty() {
		next := q.Dequeue()
		arr = append(arr, next)
		if next.left != nil {
			q.Enqueue(*next.left)
		}
		if next.right != nil {
			q.Enqueue(*next.right)
		}
	}
	return arr
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

func (t *TreeNode[K, V]) search(key K) *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if key < t.key {
		return t.left.search(key)
	} else if key > t.key {
		return t.right.search(key)
	}
	return t
}

func (t *TreeNode[K, V]) min() *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.left != nil {
		return t.left.min()
	}
	return t
}

func (t *TreeNode[K, V]) del(key K) *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if key < t.key {
		t.left = t.left.del(key)
		return t
	} else if key > t.key {
		t.right = t.right.del(key)
		return t
	}
	if t.right != nil {
		min := t.right.min()
		min.right = t.right.del(min.key)
		min.left = t.left
		t = min
		return t
	} else {
		return t.left
	}
}

func (t *TreeNode[K, V]) max() *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.right != nil {
		return t.right.max()
	}
	return t
}

func (t *TreeNode[K, V]) floor(key K) *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if key < t.key {
		return t.left.floor(key)
	} else if key > t.key {
		rightFloor := t.right.floor(key)
		if rightFloor == nil {
			return t
		}
		return rightFloor
	}
	return t
}

func (t *TreeNode[K, V]) ceil(key K) *TreeNode[K, V] {
	if t == nil {
		return nil
	}
	if key < t.key {
		leftCeil := t.left.ceil(key)
		if leftCeil == nil {
			return t
		}
		return leftCeil
	} else if key > t.key {
		return t.right.ceil(key)
	}
	return t
}

func main() {
	root := BinarySearchTree[string, string]{}
	root.Insert("V", "V")
	root.Insert("MM", "MM")
	root.Insert("UU", "UU")
	root.Insert("A", "A")
	root.Insert("PPPP", "PPPP")
	root.Insert("KK", "KK")
	root.Insert("PQ", "PQ")
	root.Insert("M", "M")
	root.Insert("ZA", "ZA")
	root.Insert("LLL", "LLL")
	root.Insert("P", "P")
	root.Insert("OOO", "OOO")
	root.Insert("WWW", "WWW")
	root.Insert("C", "C")
	root.Insert("LL", "LL")
	root.Insert("KKK", "KKK")
	root.Insert("HHH", "HHH")
	root.Insert("HHHH", "HHHH")
	root.Insert("WW", "WW")
	root.Insert("BBBB", "BBBB")
	root.Insert("OO", "OO")
	root.Insert("III", "III")
	root.Insert("TTTT", "TTTT")
	root.Insert("KKK", "KKK")
	root.Insert("I", "I")
	root.Insert("FFFFFF", "FFFFFF")
	root.Insert("B", "B")
	root.Insert("DDD", "DDD")
	root.Insert("VVVV", "VVVV")
	PrettyPrint(&root)
	root.Delete("A")
	PrettyPrint(&root)
	root.Delete("PQ")
	PrettyPrint(&root)
	root.Delete("MM")
	PrettyPrint(&root)
}
