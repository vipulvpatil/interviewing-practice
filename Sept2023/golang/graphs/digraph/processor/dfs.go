package processor

import "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"

func Dfs[T comparable](diGraphProcessor TraversalDiGraphProcessor[T]) {
	acc := accumulators.StackAccumulator[T]{}
	graph := diGraphProcessor.DiGraph()
	for k := range graph.Adjacency() {
		if !diGraphProcessor.Visited(k) {
			acc.Add(k)
			diGraphProcessor.processWithAccumulator(&acc)
		}
	}
}

func DfsFromSource[T comparable](diGraphProcessor TraversalDiGraphProcessor[T], s T) {
	acc := accumulators.StackAccumulator[T]{}
	acc.Add(s)
	diGraphProcessor.processWithAccumulator(&acc)
}
