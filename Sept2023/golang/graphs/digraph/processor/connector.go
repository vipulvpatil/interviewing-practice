package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

type StronglyConnectorDiGraphProcessor[T comparable] interface {
	DiGraphProcessor[T]
	IsConnected(v T, w T) bool
	Count() int
	ConnectionId(v T) int
}

type connectorDiGraphProcessor[T comparable] struct {
	diGraphProcessor[T]
	connections map[T]int
	count       int
}

func NewStronglyConnectorDiGraphProcessor[T comparable](g *digraph.DiGraph[T]) *connectorDiGraphProcessor[T] {
	if g == nil {
		return nil
	}

	reverseGraph := g.Reverse()
	reverseDiGraphProc := NewTraversalDiGraphProcessor[T](reverseGraph)

	acc := stack.Stack[T]{}
	reversePostOrder := stack.Stack[T]{}
	for k := range reverseGraph.Adjacency() {
		if !reverseDiGraphProc.Visited(k) {
			acc.Push(k)
			processPostOrderWithAccumulator[T](reverseDiGraphProc, &acc, &reversePostOrder)
		}
	}
	reversePop := []T{}
	for !reversePostOrder.IsEmpty() {
		reversePop = append(reversePop, *reversePostOrder.Pop())
	}

	connectorProc := connectorDiGraphProcessor[T]{
		diGraphProcessor: diGraphProcessor[T]{
			diGraph: g,
			visited: make(map[T]bool),
		},
		connections: make(map[T]int),
		count:       0,
	}

	s := accumulators.StackAccumulator[T]{}
	for _, k := range reversePop {
		if !connectorProc.Visited(k) {
			s.Add(k)
			processConnections[T](&connectorProc, &s)
		}
	}

	return &connectorProc
}

func (c *connectorDiGraphProcessor[T]) IsConnected(v T, w T) bool {
	return c.ConnectionId(v) == c.ConnectionId(w)
}
func (c *connectorDiGraphProcessor[T]) Count() int {
	return c.count
}
func (c *connectorDiGraphProcessor[T]) ConnectionId(v T) int {
	return c.connections[v]
}

func processConnections[T comparable](
	c *connectorDiGraphProcessor[T],
	acc accumulators.Accumulator[T],
) {
	for !acc.IsEmpty() {
		vertex := acc.Remove()
		v := *vertex
		if !c.Visited(v) {
			c.Visit(v)
			c.connections[v] = c.count
			graph := c.DiGraph()
			for _, a := range graph.Adj(v) {
				if !c.Visited(a) {
					acc.Add(a)
				}
			}
		}
	}
	c.count++
}
