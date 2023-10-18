package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
	"golang.org/x/exp/constraints"
)

type NodeColor int

const (
	RED   NodeColor = 0
	WHITE NodeColor = 1
)

type LLRBBinarySearchTree[K constraints.Ordered, V any] struct {
	root *LLRBTreeNode[K, V]
}

func (b *LLRBBinarySearchTree[K, V]) InOrder() []*LLRBTreeNode[K, V] {
	return b.root.inOrder()
}

func (b *LLRBBinarySearchTree[K, V]) Insert(key K, value V) {
	b.root = b.root.add(key, value)
}

func (b *LLRBBinarySearchTree[K, V]) Delete(key K) {
	b.root = b.root.del(key)
}

func (b *LLRBBinarySearchTree[K, V]) Search(key K) *LLRBTreeNode[K, V] {
	return b.root.search(key)
}

func (b *LLRBBinarySearchTree[K, V]) Min() *LLRBTreeNode[K, V] {
	return b.root.min()
}

func (b *LLRBBinarySearchTree[K, V]) Max() *LLRBTreeNode[K, V] {
	return b.root.max()
}

func (b *LLRBBinarySearchTree[K, V]) Floor(key K) *LLRBTreeNode[K, V] {
	return b.root.floor(key)
}

func (b *LLRBBinarySearchTree[K, V]) Ceil(key K) *LLRBTreeNode[K, V] {
	return b.root.ceil(key)
}

func (b *LLRBBinarySearchTree[K, V]) Print() string {
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

func PrettyPrint(b *LLRBBinarySearchTree[string, string]) {
	levelIndex := make(map[LLRBTreeNode[string, string]]int)
	blackCount := make(map[LLRBTreeNode[string, string]]int)
	levelIndex[*b.root] = 0
	blackCount[*b.root] = 0
	for _, node := range b.BFS() {
		if node.left != nil {
			levelIndex[*node.left] = levelIndex[*node] + 1
			if node.left.isRed() {
				blackCount[*node.left] = blackCount[*node]
			} else {
				blackCount[*node.left] = blackCount[*node] + 1
			}
		}
		if node.right != nil {
			levelIndex[*node.right] = levelIndex[*node] + 1
			if node.right.isRed() {
				blackCount[*node.right] = blackCount[*node]
			} else {
				blackCount[*node.right] = blackCount[*node] + 1
			}
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
		color := "B"
		if node.isRed() {
			color = "R"
		}
		fmt.Printf("%s(%s)", node.value, color)
		if node.left == nil && node.right == nil {
			fmt.Printf("[%d]", blackCount[*node])
		}
		if node.right != nil {
			if node.right.isRed() {
				fmt.Printf("**")
			}
			fmt.Printf(">%s ", node.right.value)
		}
	}
	fmt.Println()
}

func (b *LLRBBinarySearchTree[K, V]) BFS() []*LLRBTreeNode[K, V] {
	arr := []*LLRBTreeNode[K, V]{}
	q := queue.Queue[LLRBTreeNode[K, V]]{}
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

type LLRBTreeNode[K constraints.Ordered, V any] struct {
	key   K
	value V
	color NodeColor
	left  *LLRBTreeNode[K, V]
	right *LLRBTreeNode[K, V]
}

func (t *LLRBTreeNode[K, V]) inOrder() []*LLRBTreeNode[K, V] {
	arr := []*LLRBTreeNode[K, V]{}
	if t.left != nil {
		arr = append(arr, t.left.inOrder()...)
	}
	arr = append(arr, t)
	if t.right != nil {
		arr = append(arr, t.right.inOrder()...)
	}
	return arr
}

func (t *LLRBTreeNode[K, V]) add(key K, value V) *LLRBTreeNode[K, V] {
	if t == nil {
		return &LLRBTreeNode[K, V]{
			key:   key,
			value: value,
			color: RED,
		}
	}

	if key < t.key {
		t.left = t.left.add(key, value)
	} else if key > t.key {
		t.right = t.right.add(key, value)
	} else {
		t.value = value
	}

	if t.right.isRed() && !t.left.isRed() {
		t = t.rotateLeft()
	}
	if t != nil && t.left.isRed() && t.left != nil && t.left.left.isRed() {
		t = t.rotateRight()
	}
	if t != nil && t.right.isRed() && t.left.isRed() {
		t.flipColors()
	}

	return t
}

func (t *LLRBTreeNode[K, V]) flipColors() {
	t.left.color = 1 - t.left.color
	t.right.color = 1 - t.right.color
	t.color = 1 - t.color
}

func (t *LLRBTreeNode[K, V]) isRed() bool {
	if t == nil {
		return false
	}
	return t.color == RED
}

func (t *LLRBTreeNode[K, V]) rotateLeft() *LLRBTreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.right != nil {
		x := t.right
		t.right = x.left
		x.left = t
		x.color = t.color
		t.color = RED
		return x
	}
	return t
}

func (t *LLRBTreeNode[K, V]) rotateRight() *LLRBTreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.left != nil {
		x := t.left
		t.left = x.right
		x.right = t
		x.color = t.color
		t.color = RED
		return x
	}
	return t
}

func (t *LLRBTreeNode[K, V]) search(key K) *LLRBTreeNode[K, V] {
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

func (t *LLRBTreeNode[K, V]) min() *LLRBTreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.left != nil {
		return t.left.min()
	}
	return t
}

func (t *LLRBTreeNode[K, V]) del(key K) *LLRBTreeNode[K, V] {
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

func (t *LLRBTreeNode[K, V]) max() *LLRBTreeNode[K, V] {
	if t == nil {
		return nil
	}
	if t.right != nil {
		return t.right.max()
	}
	return t
}

func (t *LLRBTreeNode[K, V]) floor(key K) *LLRBTreeNode[K, V] {
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

func (t *LLRBTreeNode[K, V]) ceil(key K) *LLRBTreeNode[K, V] {
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
	root := LLRBBinarySearchTree[string, string]{}
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
	// root.Delete("A")
	// PrettyPrint(&root)
	// root.Delete("PQ")
	// PrettyPrint(&root)
	// root.Delete("MM")
	// PrettyPrint(&root)
}
