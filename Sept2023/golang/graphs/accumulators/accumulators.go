package accumulators

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/queue"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

type Accumulator[T comparable] interface {
	Add(t T)
	Remove() *T
	IsEmpty() bool
}

type StackAccumulator[T comparable] struct {
	stack.Stack[T]
}

func (sa *StackAccumulator[T]) Add(t T) {
	sa.Push(t)
}

func (sa *StackAccumulator[T]) Remove() *T {
	return sa.Pop()
}

type QueueAccumulator[T comparable] struct {
	queue.Queue[T]
}

func (sa *QueueAccumulator[T]) Add(t T) {
	sa.Enqueue(t)
}

func (sa *QueueAccumulator[T]) Remove() *T {
	return sa.Dequeue()
}
