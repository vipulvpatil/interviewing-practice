package digraph

import "fmt"

type FlowGraph[T comparable] struct {
	adjacency map[T][]*FlowEdge[T]
}

func NewFlowGraph[T comparable]() *FlowGraph[T] {
	return &FlowGraph[T]{
		adjacency: make(map[T][]*FlowEdge[T]),
	}
}

func (g *FlowGraph[T]) AddVertex(v T) {
	if _, vOk := g.adjacency[v]; !vOk {
		g.adjacency[v] = []*FlowEdge[T]{}
	}
}

func (g *FlowGraph[T]) AddFlowEdge(v, w T, capacity float64) {
	g.AddVertex(v)
	g.AddVertex(w)
	edge := NewFlowEdge[T](v, w, 0, capacity)
	g.adjacency[v] = append(g.adjacency[v], &edge)
	g.adjacency[w] = append(g.adjacency[w], &edge)
}

func (g *FlowGraph[T]) Adj(v T) []*FlowEdge[T] {
	return g.adjacency[v]
}

func (g *FlowGraph[T]) V() int {
	return len(g.adjacency)
}

func (g *FlowGraph[T]) E() int {
	edgeCount := 0
	for _, a := range g.adjacency {
		edgeCount += len(a)
	}
	return edgeCount
}

func (g FlowGraph[T]) String() string {
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

func (g *FlowGraph[T]) Adjacency() map[T][]*FlowEdge[T] {
	return g.adjacency
}
