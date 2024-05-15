package nary

import "fmt"

type Node struct {
	Children []*Node
	Value    string
}

func PreOrderNode(node *Node, ch chan<- string) {
	defer close(ch)
	if node != nil {
		preOrderInternalNode(node, ch)
	}
}

func preOrderInternalNode(node *Node, ch chan<- string) {
	if node == nil {
		return
	}

	if len([]rune(node.Value)) > 0 {
		ch <- node.Value
	}

	for _, child := range node.Children {
		preOrderInternalNode(child, ch)
	}
}

func AreSame(node1 *Node, node2 *Node) bool {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go PreOrderNode(node1, ch1)
	go PreOrderNode(node2, ch2)

	val1, ok1 := <-ch1
	val2, ok2 := <-ch2
	for ok1 && ok2 {
		fmt.Println(val1, ok1)
		fmt.Println(val2, ok2)
		if val1 != val2 {
			return false
		}
		val1, ok1 = <-ch1
		val2, ok2 = <-ch2
	}
	if ok1 || ok2 {
		return false
	}
	return true
}
