package processor

import (
	"container/heap"
	"sort"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

func MSTWithKruskal[T comparable](g *graph.EdgeWeightedGraph[T]) []graph.Edge[T] {
	allEdges := graph.SortByWeight[T](g.Edges())
	sort.Sort(allEdges)
	mstEdges := []graph.Edge[T]{}
	mstEdgeCount := 0
	for _, edge := range allEdges {
		newGraph := graph.NewEdgeWeightedGraph[T]()
		for _, mstEdge := range mstEdges {
			newGraph.AddEdge(mstEdge)
		}
		newGraph.AddEdge(edge)
		if !CyclesDetectedInEdgeWeightedGraph(*newGraph) {
			mstEdges = append(mstEdges, edge)
			mstEdgeCount++
		}
		if mstEdgeCount >= g.V()-1 {
			break
		}
	}
	return mstEdges
}

func MSTWithPrim[T comparable](g *graph.EdgeWeightedGraph[T]) []graph.Edge[T] {
	mstEdges := []graph.Edge[T]{}
	includedVertices := make(map[T]bool)
	pq := graph.MinWeightedEdgePQ[T]{}

	firstVertex := g.Edges()[0].Either()
	includedVertices[firstVertex] = true
	for _, edge := range g.Adj(firstVertex) {
		pq.Push(edge)
	}
	heap.Init(&pq)
	vertextCount := g.V()
	for len(mstEdges) < vertextCount-1 && pq.Len() > 0 {
		edge := heap.Pop(&pq).(graph.Edge[T])
		v := edge.Either()
		_, okV := includedVertices[v]
		if !okV {
			v = edge.Other(v)
		}
		w := edge.Other(v)
		_, okW := includedVertices[w]
		if !okW {
			mstEdges = append(mstEdges, edge)
			includedVertices[w] = true
			for _, edge := range g.Adj(w) {
				heap.Push(&pq, edge)
			}
		}
	}
	return mstEdges
}
