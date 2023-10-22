package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

func TopologicalOrder[T comparable](diGraphProcessor TraversalDiGraphProcessor[T]) []T {
	graph := diGraphProcessor.DiGraph()
	acc := stack.Stack[T]{}
	reversePostOrder := stack.Stack[T]{}
	for k := range graph.Adjacency() {
		if !diGraphProcessor.Visited(k) {
			acc.Push(k)
			processPostOrderWithAccumulator[T](diGraphProcessor, &acc, &reversePostOrder)
		}
	}
	reversePop := []T{}
	for !reversePostOrder.IsEmpty() {
		reversePop = append(reversePop, *reversePostOrder.Pop())
	}
	return reversePop
}

func processPostOrderWithAccumulator[T comparable](
	g TraversalDiGraphProcessor[T],
	acc *stack.Stack[T],
	reversePostOrder *stack.Stack[T],
) {
	for !acc.IsEmpty() {
		vertex := acc.Top()
		v := *vertex
		if !g.Visited(v) {
			g.Visit(v)
			graph := g.DiGraph()
			for _, a := range graph.Adj(v) {
				if !g.Visited(a) {
					acc.Push(a)
				}
			}
		} else {
			value := acc.Pop()
			reversePostOrder.Push(*value)
		}
	}
}
