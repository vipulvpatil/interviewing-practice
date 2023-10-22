package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

type TraversalGraphProcessor[T comparable] interface {
	GraphProcessor[T]
	Result() []T
	processWithAccumulator(acc accumulators.Accumulator[T])
}

type traversalGraphProcessor[T comparable] struct {
	graphProcessor[T]
	result []T
}

func NewTraversalGraphProcessor[T comparable](g *graph.Graph[T]) *traversalGraphProcessor[T] {
	if g == nil {
		return nil
	}
	return &traversalGraphProcessor[T]{
		graphProcessor: graphProcessor[T]{
			graph:   g,
			visited: make(map[T]bool),
		},
	}
}

func (g *traversalGraphProcessor[T]) Result() []T {
	return g.result
}

func (g *traversalGraphProcessor[T]) processWithAccumulator(acc accumulators.Accumulator[T]) {
	for !acc.IsEmpty() {
		vertex := acc.Remove()
		v := *vertex
		if !g.Visited(v) {
			g.Visit(v)
			g.result = append(g.result, v)
			graph := g.Graph()
			for _, a := range graph.Adj(v) {
				if !g.Visited(a) {
					acc.Add(a)
				}
			}
		}
	}
}
