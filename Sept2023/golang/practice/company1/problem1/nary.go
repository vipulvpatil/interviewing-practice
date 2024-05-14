package nary

type Node struct {
	Children []*Node
	Value string
}

func PreOrderNode(node *Node, ch chan<- string) {
	if node != nil {
		return
	}

	preOrderInternalNode(node, ch)

	ch.Close()
}

func preOrderInternalNode(node *Node, ch chan<- string) {
	if node != nil {
		return
	}

	if len([]rune(node.Value)) > 0 {
		ch <- node.Value
	}

	for _, child := range node.Children {
		preOrderInternalNode(child, ch)
	}
}
