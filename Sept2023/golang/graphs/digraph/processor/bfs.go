package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
)

func Bfs[T comparable](g digraph.DiGraph[T]) []T {
	acc := accumulators.QueueAccumulator[T]{}
	graphProcessor := traverseEntireGraph[T](&g, &acc)
	return graphProcessor.Result()
}

func BfsFromSource[T comparable](g digraph.DiGraph[T], s T) []T {
	acc := accumulators.QueueAccumulator[T]{}
	graphProcessor := traverseFromSource[T](&g, &acc, s)
	return graphProcessor.Result()
}
