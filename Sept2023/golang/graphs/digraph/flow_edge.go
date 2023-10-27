package digraph

type FlowEdge[T comparable] struct {
	v        T
	w        T
	flow     float64
	capacity float64
}

func NewFlowEdge[T comparable](v, w T, flow, capacity float64) FlowEdge[T] {
	return FlowEdge[T]{
		v:        v,
		w:        w,
		flow:     flow,
		capacity: capacity,
	}
}

func (e *FlowEdge[T]) Source() T {
	return e.v
}

func (e *FlowEdge[T]) Target() T {
	return e.w
}

func (e *FlowEdge[T]) Other(x T) T {
	if e.v == x {
		return e.w
	} else {
		return e.v
	}
}

func (e *FlowEdge[T]) Flow() float64 {
	return e.flow
}

func (e *FlowEdge[T]) Capacity() float64 {
	return e.capacity
}

func (e *FlowEdge[T]) ResidualCapacityTo(x T) float64 {
	if x == e.v {
		return e.flow
	}
	if x == e.w {
		return e.capacity - e.flow
	}
	return 0
}

func (e *FlowEdge[T]) AddFlowTo(x T, delta float64) {
	if x == e.v {
		e.flow = e.flow - delta
	} else if x == e.w {
		e.flow = e.flow + delta
	}
}
