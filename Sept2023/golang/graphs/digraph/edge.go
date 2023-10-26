package digraph

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

func (e *Edge[T]) Source() T {
	return e.v
}

func (e *Edge[T]) Target() T {
	return e.w
}

func (e *Edge[T]) Weight() float64 {
	return e.weight
}

type SortByWeight[T comparable] []Edge[T]

func (a SortByWeight[T]) Len() int {
	return len(a)
}
func (a SortByWeight[T]) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByWeight[T]) Less(i, j int) bool {
	return a[i].weight < a[j].weight
}

type MinWeightedEdgePQ[T comparable] []Edge[T]

func (pq MinWeightedEdgePQ[T]) Len() int {
	return len(pq)
}
func (pq MinWeightedEdgePQ[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq MinWeightedEdgePQ[T]) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq *MinWeightedEdgePQ[T]) Push(e any) {
	*pq = append(*pq, e.(Edge[T]))
}

func (pq *MinWeightedEdgePQ[T]) Pop() any {
	queue := *pq
	n := len(queue)
	item := queue[n-1]
	*pq = queue[0 : n-1]
	return item
}
