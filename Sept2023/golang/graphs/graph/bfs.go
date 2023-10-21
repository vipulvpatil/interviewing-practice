package graph

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
)

func NewBFSProcessor[T comparable](graph *Graph[T]) *GraphProcessor[T] {
	b := &GraphProcessor[T]{
		graph:   graph,
		result:  []T{},
		visited: make(map[T]bool),
	}

	queue := queue.Queue[T]{}
	for k := range graph.adjacency {
		if !b.marked(k) {
			queue.Enqueue(k)
			for !queue.IsEmpty() {
				b.bfs(&queue)
			}
		}
	}

	return b
}

func NewBFSProcessorWithSource[T comparable](graph *Graph[T], s T) *GraphProcessor[T] {
	b := &GraphProcessor[T]{
		graph:   graph,
		result:  []T{},
		visited: make(map[T]bool),
	}

	queue := queue.Queue[T]{}

	queue.Enqueue(s)
	for !queue.IsEmpty() {
		b.bfs(&queue)
	}

	return b
}

func (b *GraphProcessor[T]) bfs(queue *queue.Queue[T]) {
	v := queue.Dequeue()
	if !b.marked(*v) {
		b.visited[*v] = true
		b.result = append(b.result, *v)
		for _, a := range b.graph.adjacency[*v] {
			if !b.marked(a) {
				queue.Enqueue(a)
			}
		}
	}
}
