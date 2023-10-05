package sortable

type Sortable interface {
	Less(i, j int) bool
	Len() int
}

type SortableInPlace interface {
	Sortable
	Swap(i, j int)
}

type SortableOutOfPlace interface {
	Sortable
	Get(i int) any
	Set(i int, value any)
}

type InPlaceSorter interface {
	Sort(arr SortableInPlace)
}
type OutOfPlaceSorter interface {
	Sort(arr SortableOutOfPlace)
}
