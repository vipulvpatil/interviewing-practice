package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
)

func Dfs[T comparable](g digraph.DiGraph[T]) []T {
	acc := accumulators.StackAccumulator[T]{}
	graphProcessor := traverseEntireGraph[T](&g, &acc)
	return graphProcessor.Result()
}

func DfsFromSource[T comparable](g digraph.DiGraph[T], s T) []T {
	acc := accumulators.StackAccumulator[T]{}
	graphProcessor := traverseFromSource[T](&g, &acc, s)
	return graphProcessor.Result()
}
