package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph/processor"
)

func main() {
	connectionCheck()
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
