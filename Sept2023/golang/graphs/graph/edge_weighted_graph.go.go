package graph

import "fmt"

type EdgeWeightedGraph[T comparable] struct {
	adjacency map[T][]Edge[T]
	edges     []Edge[T]
}

func NewEdgeWeightedGraph[T comparable]() *EdgeWeightedGraph[T] {
	return &EdgeWeightedGraph[T]{
		adjacency: make(map[T][]Edge[T]),
		edges:     []Edge[T]{},
	}
}

func (g *EdgeWeightedGraph[T]) AddVertex(v T) {
	if _, vOk := g.adjacency[v]; !vOk {
		g.adjacency[v] = []Edge[T]{}
	}
}

func (g *EdgeWeightedGraph[T]) AddEdge(e Edge[T]) {
	v := e.Either()
	w := e.Other(v)
	vAdjacency, vOk := g.adjacency[v]
	wAdjacency, wOk := g.adjacency[w]
	if !vOk {
		vAdjacency = []Edge[T]{}
	}
	if !wOk {
		wAdjacency = []Edge[T]{}
	}
	g.adjacency[v] = append(vAdjacency, e)
	g.adjacency[w] = append(wAdjacency, e)
	g.edges = append(g.edges, e)
}

func (g *EdgeWeightedGraph[T]) Adj(v T) []Edge[T] {
	return g.adjacency[v]
}

func (g *EdgeWeightedGraph[T]) Edges() []Edge[T] {
	return g.edges
}

func (g *EdgeWeightedGraph[T]) V() int {
	return len(g.adjacency)
}

func (g *EdgeWeightedGraph[T]) E() int {
	doubleEdgeCount := 0
	for _, a := range g.adjacency {
		doubleEdgeCount += len(a)
	}
	return doubleEdgeCount / 2
}

func (g EdgeWeightedGraph[T]) String() string {
	retString := ""
	for v, vAdj := range g.adjacency {
		for _, w := range vAdj {
			retString = fmt.Sprintf("%s%v - %v\n", retString, v, w)
		}
	}
	return retString
}

func (g *EdgeWeightedGraph[T]) Adjacency() map[T][]Edge[T] {
	return g.adjacency
}
