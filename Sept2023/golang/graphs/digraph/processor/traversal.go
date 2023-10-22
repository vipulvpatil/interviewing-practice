package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
)

type TraversalDiGraphProcessor[T comparable] interface {
	DiGraphProcessor[T]
	Result() []T
	processWithAccumulator(acc accumulators.Accumulator[T])
}

type traversalDiGraphProcessor[T comparable] struct {
	diGraphProcessor[T]
	result []T
}

func NewTraversalDiGraphProcessor[T comparable](g *digraph.DiGraph[T]) *traversalDiGraphProcessor[T] {
	if g == nil {
		return nil
	}
	return &traversalDiGraphProcessor[T]{
		diGraphProcessor: diGraphProcessor[T]{
			diGraph: g,
			visited: make(map[T]bool),
		},
	}
}

func (g *traversalDiGraphProcessor[T]) Result() []T {
	return g.result
}

func (g *traversalDiGraphProcessor[T]) processWithAccumulator(acc accumulators.Accumulator[T]) {
	for !acc.IsEmpty() {
		vertex := acc.Remove()
		v := *vertex
		if !g.Visited(v) {
			g.Visit(v)
			g.result = append(g.result, v)
			graph := g.DiGraph()
			for _, a := range graph.Adj(v) {
				if !g.Visited(a) {
					acc.Add(a)
				}
			}
		}
	}
}
