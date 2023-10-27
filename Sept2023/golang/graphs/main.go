package main

import (
	"fmt"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	diprocessor "github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph/processor"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/graph/processor"
)

func main() {
	maxFlowMinCutCheck()
}

func maxFlowMinCutCheck() {
	g := digraph.NewFlowGraph[int]()
	g.AddFlowEdge(0, 1, 10)
	g.AddFlowEdge(0, 2, 5)
	g.AddFlowEdge(0, 3, 15)
	g.AddFlowEdge(1, 2, 4)
	g.AddFlowEdge(2, 3, 4)
	g.AddFlowEdge(3, 6, 16)
	g.AddFlowEdge(1, 4, 9)
	g.AddFlowEdge(1, 5, 15)
	g.AddFlowEdge(2, 5, 8)
	g.AddFlowEdge(4, 5, 15)
	g.AddFlowEdge(4, 7, 10)
	g.AddFlowEdge(5, 7, 10)
	g.AddFlowEdge(5, 6, 15)
	g.AddFlowEdge(6, 2, 6)
	g.AddFlowEdge(6, 7, 10)

	m := diprocessor.NewMaxFlowMinCut[int](*g, 0, 7)
	fmt.Println(m.MaxFlow())
	fmt.Println(m.MinCut())
}

func findShortestPath() {
	g := digraph.NewEdgeWeightedDiGraph[int]()
	g.AddEdge(0, 1, 5.0)
	g.AddEdge(0, 4, 9.0)
	g.AddEdge(0, 7, 8.0)
	g.AddEdge(1, 2, 12.0)
	g.AddEdge(1, 3, 15.0)
	g.AddEdge(1, 7, 4.0)
	g.AddEdge(2, 3, 3.0)
	g.AddEdge(2, 6, 11.0)
	g.AddEdge(3, 6, 9.0)
	g.AddEdge(4, 5, 4.0)
	g.AddEdge(4, 6, 20.0)
	g.AddEdge(4, 7, 5.0)
	g.AddEdge(5, 2, 1.0)
	g.AddEdge(5, 6, 13.0)
	g.AddEdge(7, 5, 6.0)
	g.AddEdge(7, 2, 7.0)

	fmt.Println(diprocessor.AllShortestPaths[int](g))
}

func findMST() {
	g := graph.NewEdgeWeightedGraph[int]()
	g.AddEdge(graph.NewEdge[int](4, 5, 0.35))
	g.AddEdge(graph.NewEdge[int](4, 7, 0.37))
	g.AddEdge(graph.NewEdge[int](5, 7, 0.28))
	g.AddEdge(graph.NewEdge[int](0, 7, 0.16))
	g.AddEdge(graph.NewEdge[int](1, 5, 0.32))
	g.AddEdge(graph.NewEdge[int](0, 4, 0.38))
	g.AddEdge(graph.NewEdge[int](2, 3, 0.17))
	g.AddEdge(graph.NewEdge[int](1, 7, 0.19))
	g.AddEdge(graph.NewEdge[int](0, 2, 0.26))
	g.AddEdge(graph.NewEdge[int](1, 2, 0.36))
	g.AddEdge(graph.NewEdge[int](1, 3, 0.29))
	g.AddEdge(graph.NewEdge[int](2, 7, 0.34))
	g.AddEdge(graph.NewEdge[int](6, 2, 0.40))
	g.AddEdge(graph.NewEdge[int](3, 6, 0.52))
	g.AddEdge(graph.NewEdge[int](6, 0, 0.58))
	g.AddEdge(graph.NewEdge[int](6, 4, 0.93))

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g.Edges())

	fmt.Println("Kruskal")
	kruskalMSTEdges := processor.MSTWithKruskal[int](g)
	kruskalMSTEdgesTotalWeight := 0.0
	for _, edge := range kruskalMSTEdges {
		kruskalMSTEdgesTotalWeight += edge.Weight()
	}
	fmt.Println(kruskalMSTEdges)
	fmt.Println(kruskalMSTEdgesTotalWeight)
	fmt.Println("Prim")
	primMSTEdges := processor.MSTWithKruskal[int](g)
	primMSTEdgesTotalWeight := 0.0
	for _, edge := range primMSTEdges {
		primMSTEdgesTotalWeight += edge.Weight()
	}
	fmt.Println(primMSTEdges)
	fmt.Println(primMSTEdgesTotalWeight)
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

	fmt.Println(diprocessor.Dfs[int](*g))
	fmt.Println(diprocessor.Bfs[int](*g))
	fmt.Println(diprocessor.DfsFromSource[int](*g, 0))
	fmt.Println(diprocessor.BfsFromSource[int](*g, 0))

	fmt.Println()
	p := diprocessor.NewTraversalDiGraphProcessor[int](g)
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

	fmt.Println(processor.Dfs[int](*g))
	fmt.Println(processor.Bfs[int](*g))
	fmt.Println(processor.DfsFromSource[int](*g, 0))
	fmt.Println(processor.BfsFromSource[int](*g, 0))
}

func cycleCheck() {
	g := graph.NewGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(4, 8)

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	fmt.Println(processor.Dfs[int](*g))
	fmt.Println(processor.Bfs[int](*g))
	fmt.Println(processor.DfsFromSource[int](*g, 0))
	fmt.Println(processor.BfsFromSource[int](*g, 0))

	fmt.Println(processor.DetectCyclesInEntireGraph[int](*g))
	g.AddEdge(3, 5)
	fmt.Println(processor.DetectCyclesInEntireGraph[int](*g))
}
