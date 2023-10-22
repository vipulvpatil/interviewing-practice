package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"

func Bfs[T comparable](diGraphProcessor TraversalDiGraphProcessor[T]) {
	acc := accumulators.QueueAccumulator[T]{}
	graph := diGraphProcessor.DiGraph()
	for k := range graph.Adjacency() {
		if !diGraphProcessor.Visited(k) {
			acc.Add(k)
			diGraphProcessor.processWithAccumulator(&acc)
		}
	}
}

func BfsFromSource[T comparable](diGraphProcessor TraversalDiGraphProcessor[T], s T) {
	acc := accumulators.QueueAccumulator[T]{}
	acc.Add(s)
	diGraphProcessor.processWithAccumulator(&acc)
}
