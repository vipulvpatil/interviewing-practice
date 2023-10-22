package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
)

type DiGraphProcessor[T comparable] interface {
	DiGraph() *digraph.DiGraph[T]
	Visited(v T) bool
	Visit(v T)
}

type diGraphProcessor[T comparable] struct {
	diGraph *digraph.DiGraph[T]
	visited map[T]bool
}

func (g *diGraphProcessor[T]) DiGraph() *digraph.DiGraph[T] {
	return g.diGraph
}

func (g *diGraphProcessor[T]) Visited(v T) bool {
	value, ok := g.visited[v]
	return ok && value
}

func (g *diGraphProcessor[T]) Visit(v T) {
	g.visited[v] = true
}
