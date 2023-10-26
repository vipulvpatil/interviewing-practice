package digraph

import "fmt"

type EdgeWeightedDiGraph[T comparable] struct {
	adjacency map[T][]Edge[T]
}

func NewEdgeWeightedDiGraph[T comparable]() *EdgeWeightedDiGraph[T] {
	return &EdgeWeightedDiGraph[T]{
		adjacency: make(map[T][]Edge[T]),
	}
}

func (g *EdgeWeightedDiGraph[T]) AddVertex(v T) {
	if _, vOk := g.adjacency[v]; !vOk {
		g.adjacency[v] = []Edge[T]{}
	}
}

func (g *EdgeWeightedDiGraph[T]) AddEdge(v, w T, weight float64) {
	edge := NewEdge[T](v, w, weight)
	g.adjacency[v] = append(g.adjacency[v], edge)
}

func (g *EdgeWeightedDiGraph[T]) Adj(v T) []Edge[T] {
	return g.adjacency[v]
}

func (g *EdgeWeightedDiGraph[T]) V() int {
	return len(g.adjacency)
}

func (g *EdgeWeightedDiGraph[T]) E() int {
	edgeCount := 0
	for _, a := range g.adjacency {
		edgeCount += len(a)
	}
	return edgeCount
}

func (g EdgeWeightedDiGraph[T]) String() string {
	retString := ""
	for v, vAdj := range g.adjacency {
		retString = fmt.Sprintf("%s%v ->", retString, v)
		for _, w := range vAdj {
			retString = fmt.Sprintf("%s %v", retString, w)
		}
		retString = fmt.Sprintf("%s\n", retString)
	}
	return retString
}

func (g *EdgeWeightedDiGraph[T]) Adjacency() map[T][]Edge[T] {
	return g.adjacency
}

func (g *EdgeWeightedDiGraph[T]) Reverse() *EdgeWeightedDiGraph[T] {
	reversedEdgeWeightedDiGraph := NewEdgeWeightedDiGraph[T]()
	for v, edges := range g.adjacency {
		reversedEdgeWeightedDiGraph.AddVertex(v)
		for _, edge := range edges {
			reversedEdgeWeightedDiGraph.AddVertex(edge.Target())
			reversedEdgeWeightedDiGraph.AddEdge(edge.Target(), edge.Source(), edge.Weight())
		}
	}
	return reversedEdgeWeightedDiGraph
}
