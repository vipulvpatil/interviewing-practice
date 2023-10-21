package graph

type GraphProcessor[T comparable] struct {
	graph   *Graph[T]
	visited map[T]bool
	result  []T
}

func (d *GraphProcessor[T]) marked(v T) bool {
	value, ok := d.visited[v]
	return ok && value
}

func (d *GraphProcessor[T]) Result() []T {
	return d.result
}
