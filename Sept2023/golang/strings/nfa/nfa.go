package nfa

import (
	"unicode/utf8"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/graphs/digraph/processor"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

type NFANode struct {
	char rune
}

type NFA struct {
	epGraph *digraph.DiGraph[int]
	nodes   []NFANode
}

func NewNFA(regex string) *NFA {
	nodes := []NFANode{}
	for _, r := range regex {
		nodes = append(nodes, NFANode{char: r})
	}
	nodes = append(nodes, NFANode{})
	epGraph := constructEPGraph(regex)
	return &NFA{
		epGraph: epGraph,
		nodes:   nodes,
	}
}

func (n *NFA) Match(str string) bool {
	currentNodeIndices := processor.DfsFromSource[int](*n.epGraph, 0)
	for _, r := range str {
		nextNodes := make(map[int]bool)
		for _, k := range currentNodeIndices {
			if n.nodes[k].char == r {
				nextNodes[k+1] = true
			}
		}
		reacheableNodes := make(map[int]bool)
		for k := range nextNodes {
			newNodes := processor.DfsFromSource[int](*n.epGraph, k)
			for _, l := range newNodes {
				reacheableNodes[l] = true
			}
		}
		currentNodeIndices = []int{}
		for k := range reacheableNodes {
			currentNodeIndices = append(currentNodeIndices, k)
		}
	}
	for _, k := range currentNodeIndices {
		if k == len(n.nodes)-1 {
			return true
		}
	}
	return false
}

func constructEPGraph(regex string) *digraph.DiGraph[int] {
	g := digraph.NewDiGraph[int]()
	stack := stack.Stack[int]{}
	len := utf8.RuneCountInString(regex)
	regexRunes := []rune(regex)
	for i, r := range regex {
		lp := i
		if r == '(' || r == '|' {
			stack.Push(i)
		} else if r == ')' {
			or := *stack.Pop()
			if regexRunes[or] == '|' {
				lp = *stack.Pop()
				g.AddEdge(lp, or+1)
				g.AddEdge(or, i)
			} else {
				lp = or
			}
		}
		if i < len-1 && regexRunes[i+1] == '*' {
			g.AddEdge(lp, i+1)
			g.AddEdge(i+1, lp)
		}

		if r == '(' || r == '*' || r == ')' {
			g.AddEdge(i, i+1)
		}
	}

	return g
}
