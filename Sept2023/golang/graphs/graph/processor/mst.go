package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"

type MSTProcessor[T comparable] interface {
	Graph() *graph.EdgeWeightedGraph[T]
	Edges() []graph.Edge[T]
}

type mstProcessor[T comparable] struct {
	inputGraph *graph.EdgeWeightedGraph[T]
}

func NewMSTProcessor[T comparable](g *graph.EdgeWeightedGraph[T]) *mstProcessor[T] {
	m := mstProcessor[T]{
		inputGraph: g,
	}
	// sort graph edges.
	// create a graph one edge at a time starting from lowest to highest
	// if cycle found, skip.
	// Do so until V-1 edges or we run out of edges.
	return &m
}

func (m *mstProcessor[T]) Graph() *graph.EdgeWeightedGraph[T] {
	return m.inputGraph
}

func (m *mstProcessor[T]) Edges() []graph.Edge[T] {
	return nil
}
