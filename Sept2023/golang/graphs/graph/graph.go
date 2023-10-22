package graph

import "fmt"

type Graph[T comparable] struct {
	adjacency map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjacency: make(map[T][]T),
	}
}

func (g *Graph[T]) AddVertex(v T) {
	if _, vOk := g.adjacency[v]; !vOk {
		g.adjacency[v] = []T{}
	}
}

func (g *Graph[T]) AddEdge(v, w T) {
	vAdjacency, vOk := g.adjacency[v]
	wAdjacency, wOk := g.adjacency[w]
	if !vOk {
		vAdjacency = []T{}
	}
	if !wOk {
		wAdjacency = []T{}
	}
	for _, va := range vAdjacency {
		if va == w {
			return
		}
	}
	for _, wa := range wAdjacency {
		if wa == v {
			return
		}
	}
	g.adjacency[v] = append(vAdjacency, w)
	g.adjacency[w] = append(wAdjacency, v)
}

func (g *Graph[T]) Adj(v T) []T {
	return g.adjacency[v]
}

func (g *Graph[T]) V() int {
	return len(g.adjacency)
}

func (g *Graph[T]) E() int {
	doubleEdgeCount := 0
	for _, a := range g.adjacency {
		doubleEdgeCount += len(a)
	}
	return doubleEdgeCount / 2
}

func (g Graph[T]) String() string {
	retString := ""
	for v, vAdj := range g.adjacency {
		for _, w := range vAdj {
			retString = fmt.Sprintf("%s%v - %v\n", retString, v, w)
		}
	}
	return retString
}

func (g *Graph[T]) Adjacency() map[T][]T {
	return g.adjacency
}
