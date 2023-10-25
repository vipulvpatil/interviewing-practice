package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"

func Bfs[T comparable](graphProcessor TraversalGraphProcessor[T]) {
	acc := accumulators.QueueAccumulator[T]{}
	graph := graphProcessor.Graph()
	for k := range graph.Adjacency() {
		if !graphProcessor.Visited(k) {
			acc.Add(k)
			graphProcessor.processTraversalWithAccumulator(&acc)
		}
	}
}

func BfsFromSource[T comparable](graphProcessor TraversalGraphProcessor[T], s T) {
	acc := accumulators.QueueAccumulator[T]{}
	acc.Add(s)
	graphProcessor.processTraversalWithAccumulator(&acc)
}
