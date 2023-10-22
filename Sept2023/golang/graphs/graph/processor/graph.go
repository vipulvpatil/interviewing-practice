package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"

type GraphProcessor[T comparable] interface {
	Graph() *graph.Graph[T]
	Visited(v T) bool
	Visit(v T)
}

type graphProcessor[T comparable] struct {
	graph   *graph.Graph[T]
	visited map[T]bool
}

func (g *graphProcessor[T]) Graph() *graph.Graph[T] {
	return g.graph
}

func (g *graphProcessor[T]) Visited(v T) bool {
	value, ok := g.visited[v]
	return ok && value
}

func (g *graphProcessor[T]) Visit(v T) {
	g.visited[v] = true
}
