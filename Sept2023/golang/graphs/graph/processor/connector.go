package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

type ConnectorGraphProcessor[T comparable] interface {
	GraphProcessor[T]
	IsConnected(v T, w T) bool
	Count() int
	connectionId(v T) int
}

type connectorGraphProcessor[T comparable] struct {
	graphProcessor[T]
	connections map[T]int
	count       int
}

func NewConnectorGraphProcessor[T comparable](g *graph.Graph[T]) *connectorGraphProcessor[T] {
	if g == nil {
		return nil
	}

	acc := accumulators.StackAccumulator[T]{}
	connectorProcessor := connectorGraphProcessor[T]{
		graphProcessor: graphProcessor[T]{
			graph:   g,
			visited: make(map[T]bool),
		},
		connections: make(map[T]int),
		count:       0,
	}
	for k := range g.Adjacency() {
		if !connectorProcessor.graphProcessor.Visited(k) {
			acc.Add(k)
			connectorProcessor.processWithAccumulator(&acc)
		}
	}

	return &connectorProcessor
}

func (p *connectorGraphProcessor[T]) IsConnected(v T, w T) bool {
	return p.connectionId(v) == p.connectionId(w)
}
func (p *connectorGraphProcessor[T]) Count() int {
	return p.count
}
func (p *connectorGraphProcessor[T]) connectionId(v T) int {
	return p.connections[v]
}

func (p *connectorGraphProcessor[T]) processWithAccumulator(acc accumulators.Accumulator[T]) {
	for !acc.IsEmpty() {
		vertex := acc.Remove()
		v := *vertex
		if !p.Visited(v) {
			p.Visit(v)
			p.connections[v] = p.count
			graph := p.Graph()
			for _, a := range graph.Adj(v) {
				if !p.Visited(a) {
					acc.Add(a)
				}
			}
		}
	}
	p.count++
}
