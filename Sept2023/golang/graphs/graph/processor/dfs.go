package processor

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/accumulators"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

func Dfs[T comparable](g graph.Graph[T]) []T {
	acc := accumulators.StackAccumulator[T]{}
	graphProcessor := traverseEntireGraph[T](&g, &acc)
	return graphProcessor.Result()
}

func DfsFromSource[T comparable](g graph.Graph[T], s T) []T {
	acc := accumulators.StackAccumulator[T]{}
	graphProcessor := traverseFromSource[T](&g, &acc, s)
	return graphProcessor.Result()
}
