package digraph

import "fmt"

type DiGraph[T comparable] struct {
	adjacency map[T][]T
}

func NewDiGraph[T comparable]() *DiGraph[T] {
	return &DiGraph[T]{
		adjacency: make(map[T][]T),
	}
}

func (g *DiGraph[T]) AddVertex(v T) {
	if _, vOk := g.adjacency[v]; !vOk {
		g.adjacency[v] = []T{}
	}
}

func (g *DiGraph[T]) AddEdge(v, w T) {
	vAdjacency, vOk := g.adjacency[v]
	if !vOk {
		vAdjacency = []T{}
	}
	for _, va := range vAdjacency {
		if va == w {
			return
		}
	}
	g.adjacency[v] = append(vAdjacency, w)
}

func (g *DiGraph[T]) Adj(v T) []T {
	return g.adjacency[v]
}

func (g *DiGraph[T]) V() int {
	return len(g.adjacency)
}

func (g *DiGraph[T]) E() int {
	edgeCount := 0
	for _, a := range g.adjacency {
		edgeCount += len(a)
	}
	return edgeCount
}

func (g DiGraph[T]) String() string {
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

func (g *DiGraph[T]) Adjacency() map[T][]T {
	return g.adjacency
}

func (g *DiGraph[T]) Reverse() *DiGraph[T] {
	reversedDiGraph := NewDiGraph[T]()
	for v, vAdj := range g.adjacency {
		reversedDiGraph.AddVertex(v)
		for _, w := range vAdj {
			reversedDiGraph.AddEdge(w, v)
		}
	}
	return reversedDiGraph
}
