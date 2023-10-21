package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
)

func main() {
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

	fmt.Println(graph.NewDFSProcessor[int](g).Result())
	fmt.Println(graph.NewBFSProcessor[int](g).Result())
	fmt.Println(graph.NewDFSProcessorWithSource[int](g, 0).Result())
	fmt.Println(graph.NewBFSProcessorWithSource[int](g, 0).Result())
}
