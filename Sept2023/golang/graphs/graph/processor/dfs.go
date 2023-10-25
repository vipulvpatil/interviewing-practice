package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"

func Dfs[T comparable](graphProcessor TraversalGraphProcessor[T]) {
	acc := accumulators.StackAccumulator[T]{}
	graph := graphProcessor.Graph()
	for k := range graph.Adjacency() {
		if !graphProcessor.Visited(k) {
			acc.Add(k)
			graphProcessor.processTraversalWithAccumulator(&acc)
		}
	}
}

func DfsFromSource[T comparable](graphProcessor TraversalGraphProcessor[T], s T) {
	acc := accumulators.StackAccumulator[T]{}
	acc.Add(s)
	graphProcessor.processTraversalWithAccumulator(&acc)
}
