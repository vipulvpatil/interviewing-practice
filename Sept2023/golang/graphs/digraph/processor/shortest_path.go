package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

func AllShortestPaths[T comparable](g *digraph.EdgeWeightedDiGraph[T]) map[T]map[T][]digraph.Edge[T] {
	shortestPaths := make(map[T]map[T][]digraph.Edge[T])
	for v := range g.Adjacency() {
		shortestPaths[v] = ShortestPathsFrom[T](g, v)
	}
	return shortestPaths
}

func ShortestPathsFrom[T comparable](g *digraph.EdgeWeightedDiGraph[T], source T) map[T][]digraph.Edge[T] {
	dists := make(map[T]float64)
	pathTo := make(map[T]digraph.Edge[T])

	dists[source] = 0
	q := queue.Queue[T]{}
	q.Enqueue(source)
	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, edge := range g.Adj(*v) {
			if relaxEdge[T](edge, dists, pathTo) {
				q.Enqueue(edge.Target())
			}
		}
	}

	shortestPaths := make(map[T][]digraph.Edge[T])
	for v := range g.Adjacency() {
		if _, ok := pathTo[v]; ok {
			shortestPaths[v] = calculatePath(source, v, pathTo)
		}
	}

	return shortestPaths
}

func ShortestPathFromSourceToTarget[T comparable](g *digraph.EdgeWeightedDiGraph[T], source, target T) []digraph.Edge[T] {
	pathsFromSource := ShortestPathsFrom[T](g, source)

	return pathsFromSource[target]
}

func calculatePath[T comparable](source, target T, pathTo map[T]digraph.Edge[T]) []digraph.Edge[T] {
	st := stack.Stack[digraph.Edge[T]]{}
	edge, ok := pathTo[target]
	for ok {
		st.Push(edge)
		newTarget := edge.Source()
		if newTarget == source {
			break
		}
		edge, ok = pathTo[newTarget]
	}
	if edge.Source() != source {
		return nil
	}
	path := []digraph.Edge[T]{}
	for !st.IsEmpty() {
		path = append(path, *st.Pop())
	}
	return path
}

func relaxEdge[T comparable](edge digraph.Edge[T], dists map[T]float64, pathTo map[T]digraph.Edge[T]) bool {
	target := edge.Target()
	currentTargetDist, ok := dists[target]
	newDist := dists[edge.Source()] + edge.Weight()
	if !ok {
		dists[target] = newDist
		pathTo[target] = edge
		return true
	} else {
		if currentTargetDist > newDist {
			dists[target] = newDist
			pathTo[target] = edge
			return true
		}
	}
	return false
}
