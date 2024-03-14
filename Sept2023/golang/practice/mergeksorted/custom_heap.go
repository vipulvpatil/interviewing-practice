package mergeksorted

import "sort"

type customHeapInterface interface {
	sort.Interface
	Push(x any)
	Pop() any
}

func Push(container customHeapInterface, x any) {
	container.Push(x)
	l := container.Len()
	swim(container, l-1)
}

func Pop(container customHeapInterface) any {
	l := container.Len()
	container.Swap(0, l-1)
	result := container.Pop()
	sink(container, 0)

	return result
}

func swim(container customHeapInterface, i int) {
	if i == 0 {
		return
	}
	p := (i - 1) / 2
	if container.Less(i, p) {
		container.Swap(p, i)
		swim(container, p)
	}
}

func sink(container customHeapInterface, i int) {
	l := container.Len()
	if i >= l {
		return
	}
	c1 := 2*i + 1
	c2 := 2*i + 2
	if c1 >= l {
		return
	}
	smallestC := c1

	if c2 < l {
		if container.Less(c2, smallestC) {
			smallestC = c2
		}
	}

	if container.Less(smallestC, i) {
		container.Swap(smallestC, i)
		sink(container, smallestC)
	}
}
