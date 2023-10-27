package processor

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
)

type MaxFlowMinCut[T comparable] interface {
	MaxFlow() float64
	MinCut() []T
}
type maxFlowMinCut[T comparable] struct {
	visited map[T]bool
	pathTo  map[T]*digraph.FlowEdge[T]
	maxFlow float64
	minCut  []T
}

func NewMaxFlowMinCut[T comparable](g digraph.FlowGraph[T], source, target T) MaxFlowMinCut[T] {
	m := &maxFlowMinCut[T]{}

	for m.NextAugmentedPath(&g, source, target) {
		bottleNeck := 0.0
		for v := target; v != source; {
			edge := m.pathTo[v]
			if v == target {
				bottleNeck = edge.ResidualCapacityTo(v)
			} else {
				if bottleNeck > edge.ResidualCapacityTo(v) {
					bottleNeck = edge.ResidualCapacityTo(v)
				}
			}
			v = edge.Other(v)
		}

		for v := target; v != source; {
			edge := m.pathTo[v]
			edge.AddFlowTo(v, bottleNeck)
			v = edge.Other(v)
		}
	}

	printAllEdges[T](g, source)

	m.minCut = []T{}
	for v, ok := range m.visited {
		if ok {
			m.minCut = append(m.minCut, v)
		}
	}
	m.maxFlow = 0
	for _, edge := range g.Adj(target) {
		m.maxFlow += edge.Flow()
	}

	return m
}

func (m *maxFlowMinCut[T]) NextAugmentedPath(g *digraph.FlowGraph[T], source, target T) bool {
	m.visited = make(map[T]bool)
	m.pathTo = make(map[T]*digraph.FlowEdge[T])
	q := queue.Queue[T]{}
	q.Enqueue(source)
	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, edge := range g.Adj(*v) {
			w := edge.Other(*v)
			visit, ok := m.visited[w]
			if edge.ResidualCapacityTo(w) > 0 && (!ok || !visit) {
				m.visited[w] = true
				m.pathTo[w] = edge
				q.Enqueue(w)
			}
		}
	}
	_, ok := m.visited[target]
	return ok
}

func (m *maxFlowMinCut[T]) MaxFlow() float64 {
	return m.maxFlow
}

func (m *maxFlowMinCut[T]) MinCut() []T {
	return m.minCut
}

func printAllEdges[T comparable](g digraph.FlowGraph[T], source T) {
	visited := make(map[T]bool)
	printed := make(map[digraph.FlowEdge[T]]bool)
	q := queue.Queue[T]{}
	q.Enqueue(source)
	for !q.IsEmpty() {
		v := q.Dequeue()
		visit, ok := visited[*v]
		if !ok || !visit {
			for _, edge := range g.Adj(*v) {
				w := edge.Other(*v)
				_, ok1 := printed[*edge]
				if !ok1 {
					printed[*edge] = true
					fmt.Printf("edge: %v-> %v, flow->%d, capacity->%d\n", edge.Source(), edge.Target(), int(edge.Flow()), int(edge.Capacity()))
				}
				q.Enqueue(w)
			}
			visited[*v] = true
		}
	}
}
