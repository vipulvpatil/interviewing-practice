package graph

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"

func NewDFSProcessor[T comparable](graph *Graph[T]) *GraphProcessor[T] {
	d := &GraphProcessor[T]{
		graph:   graph,
		result:  []T{},
		visited: make(map[T]bool),
	}

	stack := stack.Stack[T]{}
	for k := range graph.adjacency {
		if !d.marked(k) {
			stack.Push(k)
			for !stack.IsEmpty() {
				d.dfs(&stack)
			}
		}
	}

	return d
}

func NewDFSProcessorWithSource[T comparable](graph *Graph[T], s T) *GraphProcessor[T] {
	d := &GraphProcessor[T]{
		graph:   graph,
		result:  []T{},
		visited: make(map[T]bool),
	}

	stack := stack.Stack[T]{}

	stack.Push(s)
	for !stack.IsEmpty() {
		d.dfs(&stack)
	}

	return d
}

func (d *GraphProcessor[T]) dfs(stack *stack.Stack[T]) {
	v := stack.Pop()
	if !d.marked(*v) {
		d.visited[*v] = true
		d.result = append(d.result, *v)
		for _, a := range d.graph.adjacency[*v] {
			if !d.marked(a) {
				stack.Push(a)
			}
		}
	}
}
