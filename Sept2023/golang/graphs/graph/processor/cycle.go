package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

func DetectCyclesInEntireGraph[T comparable](g graph.Graph[T]) bool {
	visited := make(map[T]bool)
	visitedInCurrentPath := make(map[T]bool)

	for v := range g.Adjacency() {
		cycleFound := dfsRecursive[T](g, v, nil, visited, visitedInCurrentPath)
		if cycleFound {
			return true
		}
	}
	return false
}

func DetectCyclesFromSource[T comparable](g graph.Graph[T], s T) bool {
	visited := make(map[T]bool)
	visitedInCurrentPath := make(map[T]bool)
	return dfsRecursive[T](g, s, nil, visited, visitedInCurrentPath)
}

func dfsRecursive[T comparable](g graph.Graph[T], v T, parent *T, visited, visitedInCurrentPath map[T]bool) bool {
	vis, ok := visited[v]
	if ok && vis {
		vic, ok := visitedInCurrentPath[v]
		if ok {
			return vic
		}
	} else {
		visited[v] = true
		visitedInCurrentPath[v] = true
		for _, w := range g.Adj(v) {
			if parent == nil || w != *parent {
				cycleFound := dfsRecursive[T](g, w, &v, visited, visitedInCurrentPath)
				if cycleFound {
					return true
				}
			}
		}
		visitedInCurrentPath[v] = false
	}
	return false
}

func CyclesDetectedInEdgeWeightedGraph[T comparable](g graph.EdgeWeightedGraph[T]) bool {
	visited := make(map[T]bool)
	visitedInCurrentPath := make(map[T]bool)

	for v := range g.Adjacency() {
		cycleFound := dfsRecursiveInEdgeWeightedGraph[T](g, v, nil, visited, visitedInCurrentPath)
		if cycleFound {
			return true
		}
	}
	return false
}

func dfsRecursiveInEdgeWeightedGraph[T comparable](g graph.EdgeWeightedGraph[T], v T, parent *T, visited, visitedInCurrentPath map[T]bool) bool {
	vis, ok := visited[v]
	if ok && vis {
		vic, ok := visitedInCurrentPath[v]
		if ok {
			return vic
		}
	} else {
		visited[v] = true
		visitedInCurrentPath[v] = true
		for _, edge := range g.Adj(v) {
			w := edge.Other(v)
			if parent == nil || w != *parent {
				cycleFound := dfsRecursiveInEdgeWeightedGraph[T](g, w, &v, visited, visitedInCurrentPath)
				if cycleFound {
					return true
				}
			}
		}
		visitedInCurrentPath[v] = false
	}
	return false
}
