package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	diprocessor "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph/processor"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph/processor"
)

func main() {
	stronglyConnectedCheck()
}

func stronglyConnectedCheck() {
	g := digraph.NewDiGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 2)
	g.AddEdge(3, 5)
	g.AddEdge(4, 2)
	g.AddEdge(4, 3)
	g.AddEdge(5, 4)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(6, 8)
	g.AddEdge(6, 9)
	g.AddEdge(7, 6)
	g.AddEdge(7, 9)
	g.AddEdge(8, 6)
	g.AddEdge(9, 10)
	g.AddEdge(9, 11)
	g.AddEdge(10, 12)
	g.AddEdge(11, 4)
	g.AddEdge(11, 12)
	g.AddEdge(12, 9)

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	c := diprocessor.NewStronglyConnectorDiGraphProcessor[int](g)
	fmt.Println(c.Count())
	fmt.Println(c.IsConnected(0, 1), "should be false")
	fmt.Println(c.IsConnected(0, 2), "should be true")
	fmt.Println(c.IsConnected(0, 3), "should be true")
	fmt.Println(c.IsConnected(0, 4), "should be true")
	fmt.Println(c.IsConnected(0, 5), "should be true")
	fmt.Println(c.IsConnected(9, 10), "should be true")
	fmt.Println(c.IsConnected(11, 10), "should be true")
	fmt.Println(c.IsConnected(12, 10), "should be true")
	fmt.Println(c.IsConnected(6, 8), "should be true")
	fmt.Println(c.IsConnected(7, 8), "should be false")
	fmt.Println(c.IsConnected(7, 9), "should be false")

	fmt.Println(c.ConnectionId(0))
	fmt.Println(c.ConnectionId(1))
	fmt.Println(c.ConnectionId(2))
	fmt.Println(c.ConnectionId(3))
	fmt.Println(c.ConnectionId(4))
	fmt.Println(c.ConnectionId(5))
	fmt.Println(c.ConnectionId(6))
	fmt.Println(c.ConnectionId(7))
	fmt.Println(c.ConnectionId(8))
	fmt.Println(c.ConnectionId(9))
	fmt.Println(c.ConnectionId(10))
	fmt.Println(c.ConnectionId(11))
	fmt.Println(c.ConnectionId(12))
}

func digraphCheck() {
	g := digraph.NewDiGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(9, 10)
	g.AddEdge(11, 12)
	g.AddEdge(11, 3)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(3, 9)
	g.AddEdge(5, 6)
	g.AddEdge(9, 7)

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)
	fmt.Println(g.Reverse())

	var p diprocessor.TraversalDiGraphProcessor[int]

	p = diprocessor.NewTraversalDiGraphProcessor[int](g)
	diprocessor.Dfs[int](p)
	fmt.Println(p.Result())

	p = diprocessor.NewTraversalDiGraphProcessor[int](g)
	diprocessor.Bfs[int](p)
	fmt.Println(p.Result())

	p = diprocessor.NewTraversalDiGraphProcessor[int](g)
	diprocessor.DfsFromSource[int](p, 0)
	fmt.Println(p.Result())

	p = diprocessor.NewTraversalDiGraphProcessor[int](g)
	diprocessor.BfsFromSource[int](p, 0)
	fmt.Println(p.Result())

	fmt.Println()
	p = diprocessor.NewTraversalDiGraphProcessor[int](g)
	fmt.Println(diprocessor.TopologicalOrder[int](p))
}

func connectionCheck() {
	g := graph.NewGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddVertex(5)
	g.AddEdge(6, 7)
	g.AddEdge(7, 5)
	g.AddEdge(7, 8)
	g.AddEdge(9, 10)
	g.AddEdge(11, 12)
	g.AddEdge(11, 3)

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	c := processor.NewConnectorGraphProcessor[int](g)
	fmt.Println(c.Count())
	fmt.Println(c.IsConnected(0, 12), "should be true")
	fmt.Println(c.IsConnected(2, 6), "should be false")
	fmt.Println(c.IsConnected(5, 6), "should be true")
	fmt.Println(c.IsConnected(9, 10), "should be true")
	fmt.Println(c.IsConnected(8, 9), "should be false")
	fmt.Println(c.IsConnected(10, 11), "should be false")
}

func traversalCheck() {
	g := graph.NewGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(2, 5)
	g.AddEdge(5, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(9, 10)
	g.AddEdge(4, 9)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(3, 10)

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	var p processor.TraversalGraphProcessor[int]

	p = processor.NewTraversalGraphProcessor[int](g)
	processor.Dfs[int](p)
	fmt.Println(p.Result())

	p = processor.NewTraversalGraphProcessor[int](g)
	processor.Bfs[int](p)
	fmt.Println(p.Result())

	p = processor.NewTraversalGraphProcessor[int](g)
	processor.DfsFromSource[int](p, 0)
	fmt.Println(p.Result())

	p = processor.NewTraversalGraphProcessor[int](g)
	processor.BfsFromSource[int](p, 0)
	fmt.Println(p.Result())
}
