package graph

type Edge[T comparable] struct {
	v      T
	w      T
	weight float64
}

func NewEdge[T comparable](v, w T, weight float64) Edge[T] {
	return Edge[T]{
		v:      v,
		w:      w,
		weight: weight,
	}
}

func (e *Edge[T]) Either() T {
	return e.v
}

func (e *Edge[T]) Other(v T) T {
	if e.v == v {
		return e.w
	}
	return e.v
}

func (e *Edge[T]) Weight() float64 {
	return e.weight
}
